package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Medvedevsky/simple-web-application/internal/app/apiserver"
)

// путь к файлу конфигурации задаем в качетсве флага при запуске бинарника
var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	// конфигурация веб сервера
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.NewServer(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
