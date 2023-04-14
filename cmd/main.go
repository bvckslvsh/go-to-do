package main

import (
	gotodo "github.com/bvckslvsh/go-to-do"
	"github.com/bvckslvsh/go-to-do/pkg/handler"
	"github.com/bvckslvsh/go-to-do/pkg/repository"
	"github.com/bvckslvsh/go-to-do/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initializing config files: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gotodo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error while running : %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
