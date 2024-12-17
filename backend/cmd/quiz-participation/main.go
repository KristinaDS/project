package main

import (
	"backend/internal/quiz-participation/api"
	"backend/internal/quiz-participation/config"
	"backend/internal/quiz-participation/provider"
	"backend/internal/quiz-participation/usecase"
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
	use := usecase.NewUsecase(prv, cfg.QuizService.BaseURL)
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxMessageSize, use)

	srv.Run()
}
