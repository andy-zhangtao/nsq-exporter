package main

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var ENV_LIST = []string{
	"NSQD_ADDR",
	"NSQD_TOPIC",
	"NSQD_CHANNEL",
}

var _url = os.Getenv("NSQD_ADDR")

const ModuleName = "NSQD-Exporter"

func main() {

	if err := check(); err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)

	exporter()
}

func check() error {

	for _, e := range ENV_LIST {
		if os.Getenv(e) == "" {
			return errors.New(fmt.Sprint("[%s]Empty!", e))
		}
	}

	return nil
}
