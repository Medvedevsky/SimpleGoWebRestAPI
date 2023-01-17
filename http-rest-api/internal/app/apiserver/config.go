package apiserver

type Config struct {
	WebAddress  string `toml:"web_address"` // адрес на котором запускается веб сервер
	LogLevel    string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		WebAddress: ":8080",
		LogLevel:   "debug",
	}
}
