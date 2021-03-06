package app

import (
	"github.com/IamStubborN/calendar/notifier/config"
	"github.com/sirupsen/logrus"
)

func initializeConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalln(err)
	}

	return cfg
}
