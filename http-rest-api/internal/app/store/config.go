package store

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "host=localhost port=5432 user=postgres password=andrey5522 dbname=restapi_dev sslmode=disable",
	}
}
