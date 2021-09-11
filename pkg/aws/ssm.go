package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var cli *ssm.Client

func init() {
	cli = ssm.NewFromConfig(Cfg)
}

func GetParameter(ctx context.Context, parameterName string, withDecryption bool) (*string, error) {

	getParameterOutput, err := cli.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           &parameterName,
		WithDecryption: withDecryption,
	})
	if err != nil {
		return nil, err
	}
	return getParameterOutput.Parameter.Value, nil
}
