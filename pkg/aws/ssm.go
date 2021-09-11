package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	log "github.com/sirupsen/logrus"
)

func GetParameter(ctx context.Context, parameterName string) *string {
	svc := ssm.NewFromConfig(Cfg)

	getParameterOutput, err := svc.GetParameter(ctx, &ssm.GetParameterInput{Name: &parameterName})
	if err != nil {
		log.WithError(err).Print("Could not fetch SSM parameter.")
		return aws.String("")
	}
	return getParameterOutput.Parameter.Value
}
