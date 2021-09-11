package main

import (
	"context"
	"flag"
	"github.com/bpalermo/aws-ssm-config-downloader/pkg/aws"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	logLevel := os.Getenv("LOG_LEVEL")
	if strings.EqualFold(logLevel, "debug") {
		log.SetLevel(log.DebugLevel)
	} else if strings.EqualFold(logLevel, "warn") {
		log.SetLevel(log.WarnLevel)
	}
}

const (
	defaultPerm os.FileMode = 0644
)

var (
	configPath     string
	configFileName string
	parameterName  string
	withDecryption bool
)

func main() {

	flag.StringVar(&configPath, "configPath", "", "Config file location.")
	flag.StringVar(&configFileName, "configFileName", "", "Config filename.")
	flag.StringVar(&parameterName, "parameterName", "", "AWS SSM parameter name.")
	flag.BoolVar(&withDecryption, "withDecryption", false, "Decrypts values for secure string parameters.")

	flag.Parse()

	if configPath == "" {
		log.Fatal("Config path is required.")
	}

	if parameterName == "" {
		log.Fatal("Parameter name is required.")
	}

	ctx := context.TODO()
	config, err := aws.GetParameter(ctx, parameterName, withDecryption)
	if err != nil {
		log.WithError(err).Fatalf("Could not fetch config from %s.", parameterName)
	}

	log.Debug("Writing file.")
	err = os.MkdirAll(configPath, defaultPerm)
	if err != nil {
		log.WithError(err).Fatalf("Could not create path: '%s'.", configPath)
	}

	filePath := filepath.Join(configPath, configFileName)
	err = ioutil.WriteFile(filePath, []byte(*config), defaultPerm)
	if err != nil {
		log.WithError(err).Fatalf("Could not write config: '%s'.", filePath)
	}
	log.Printf("Config written to: %s.", filePath)
}
