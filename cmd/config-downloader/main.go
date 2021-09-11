package main

import (
	"context"
	"flag"
	"github.com/bpalermo/aws-ssm-config-downloader/pkg/aws"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	logLevel := os.Getenv("LOG_LEVEL")
	if strings.EqualFold(logLevel, "debug") {
		log.SetLevel(log.DebugLevel)
	} else if strings.EqualFold(logLevel, "warn") {
		log.SetLevel(log.WarnLevel)
	}
}

func main() {

	var configPath string
	var parameterName string

	flag.StringVar(&configPath, "configPath", "", "Config file location to be stored")
	flag.StringVar(&parameterName, "parameterName", "", "AWS SSM parameter name")

	flag.Parse()

	if configPath == "" {
		log.Fatal("Config path is required.")
	}

	if parameterName == "" {
		log.Fatal("Parameter name is required.")
	}

	ctx := context.TODO()
	config, err := aws.GetParameter(ctx, parameterName)
	if err != nil {
		log.WithError(err).Fatalf("Could not fetch config from %s.", parameterName)
	}

	err = ioutil.WriteFile(configPath, []byte(*config), 0644)
	if err != nil {
		log.WithError(err).Fatalf("Could not write config '%s'.", configPath)
	}
	log.Printf("Saved config to '%s'.", configPath)
}
