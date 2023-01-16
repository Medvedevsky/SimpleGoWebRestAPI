package apiserver

import "github.com/Medvedevsky/simple-web-application/internal/app/store"

type Config struct {
	WebAddress string `toml:"web_address"` // адрес на котором запускается веб сервер
	LogLevel   string `toml:"log_level"`
	Store      *store.Config
}

func NewConfig() *Config {
	return &Config{
		WebAddress: ":8080",
		LogLevel:   "debug",
		Store:      store.NewConfig(),
	}
}
