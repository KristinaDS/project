package main

import (
	"backend/internal/quiz/api"
	"backend/internal/quiz/config"
	"backend/internal/quiz/provider"
	"backend/internal/quiz/usecase"
	"flag"
	"log"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "./configs/config.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(prv)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}
