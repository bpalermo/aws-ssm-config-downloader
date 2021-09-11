package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	log "github.com/sirupsen/logrus"
)

var cli *ssm.Client

func init() {
	cli = ssm.NewFromConfig(Cfg)
}

func GetParameter(ctx context.Context, parameterName string, withDecryption bool) (*string, error) {
	log.Debugf("Fetching SSM parameter: %s", parameterName)
	if withDecryption {
		log.Debug("Parameter will be decrypted!")
	}

	getParameterOutput, err := cli.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           &parameterName,
		WithDecryption: withDecryption,
	})
	if err != nil {
		return nil, err
	}
	return getParameterOutput.Parameter.Value, nil
}
