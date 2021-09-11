package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	log "github.com/sirupsen/logrus"
)

var Cfg aws.Config

func init() {
	ctx := context.Background()

	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	var err error
	Cfg, err = config.LoadDefaultConfig(ctx)
	if err != nil {
		log.WithError(err).Print("Unable to load SDK config.")
	}
}
