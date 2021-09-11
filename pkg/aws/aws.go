package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	awsRegionKey = "AWS_REGION"
)

var Cfg aws.Config

func init() {
	region := os.Getenv(awsRegionKey)
	if region == "" {
		region = "us-east-1"
	}

	ctx := context.Background()

	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	var err error
	Cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		log.WithError(err).Print("Unable to load SDK config.")
	}
}
