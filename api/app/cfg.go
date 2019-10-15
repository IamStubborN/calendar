package app

import (
	"github.com/IamStubborN/calendar/api/config"
	"github.com/sirupsen/logrus"
)

func initializeConfig() *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalln(err)
	}

	return cfg
}
