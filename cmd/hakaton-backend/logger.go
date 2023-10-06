package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger(level log.Level) *log.Logger {
	logger := log.New()
	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)

	return logger
}