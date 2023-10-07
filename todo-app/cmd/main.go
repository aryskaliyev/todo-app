package main

import (
	"lincoln.boris/todo"
	"lincoln.boris/todo/pkg/handler"
	"lincoln.boris/todo/pkg/repository"
	"lincoln.boris/todo/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
