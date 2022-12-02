package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/daniel-888/BloXroute-CS/server/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	// open a file
	f, err := os.OpenFile("testlogrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	logrus.SetOutput(io.MultiWriter(f, os.Stdout))
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.WithField("Title", "Start of Server Logging").Info(time.Now())
	cli := cmd.NewCLI()
	if err := cli.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	logrus.WithField("Title", "The end of main function").Info(time.Now())
}
