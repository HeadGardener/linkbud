package main

import (
	"github.com/HeadGardener/linkbud/configs"
	app "github.com/HeadGardener/linkbud/internal/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initializing config: %s", err.Error())
	}

	dbconf := configs.DBConfig{}
	dbconf.Init()

	application, err := app.New(dbconf)
	if err != nil {
		logrus.Fatalf("error while creating app: %s", err.Error())
	}

	srvconf := configs.ServerConfig{}
	srvconf.Init()

	if err := application.Start(srvconf); err != nil {
		logrus.Fatalf("error while starting the app: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
