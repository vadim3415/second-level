package main

import (
	"context"
	"secondlevel/develop/dev11/internal/handler"
	"secondlevel/develop/dev11/internal/repository"
	"secondlevel/develop/dev11/internal/server"
	"secondlevel/develop/dev11/internal/service"

	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)
func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil{
			logrus.Fatalf(err.Error())
		}
	}()

	logrus.Printf("dev11 Started, port: %s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("dev11 Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("develop/dev11/.")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}