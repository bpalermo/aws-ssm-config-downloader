package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func GetParameter(ctx context.Context, parameterName string) (*string, error) {
	svc := ssm.NewFromConfig(Cfg)

	getParameterOutput, err := svc.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           &parameterName,
		WithDecryption: true,
	})
	if err != nil {
		return aws.String(""), err
	}
	return getParameterOutput.Parameter.Value, nil
}
