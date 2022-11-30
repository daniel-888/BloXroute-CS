package main

import (
	"os"

	"github.com/daniel-888/BloXroute-CS/server/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	cli := cmd.NewCLI()
	if err := cli.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}