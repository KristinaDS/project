package config

type QuizServiceConfig struct {
	BaseURL string `yaml:"base_url"`
}

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port_participation"`

	API     api     `yaml:"api"`
	Usecase usecase `yaml:"usecase"`
	DB      db      `yaml:"db"`

	QuizService QuizServiceConfig `yaml:"quiz_service"`
}

type api struct {
	MaxMessageSize int `yaml:"max_message_size"`
}

type usecase struct {
	DefaultMessage int `yaml:"default_message"`
}

type db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}
