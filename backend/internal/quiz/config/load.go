package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadConfig читает конфигурацию из YAML-файла и возвращает структуру Config
func LoadConfig(pathToFile string) (*Config, error) {
	// Преобразуем относительный путь в абсолютный
	filename, err := filepath.Abs(pathToFile)
	if err != nil {
		return nil, err
	}

	// Читаем файл конфигурации
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Декодируем содержимое YAML в структуру Config
	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}

	// Возвращаем загруженную структуру конфигурации
	return &cfg, nil
}
