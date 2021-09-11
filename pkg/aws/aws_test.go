package aws

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	awsAccessKeyIdKey     = "AWS_ACCESS_KEY_ID"
	awsSecretAccessKeyKey = "AWS_SECRET_ACCESS_KEY"
)

func TestSuccess(t *testing.T) {
	region := "us-east-1"
	t.Setenv(awsRegionKey, region)
	t.Setenv(awsAccessKeyIdKey, "accessKey")
	t.Setenv(awsSecretAccessKeyKey, "secret")
	assert.NotNil(t, Cfg)
	assert.Equal(t, Cfg.Region, region)
}
