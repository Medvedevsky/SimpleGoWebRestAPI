package apiserver

import (
	"io"
	"net/http"

	"github.com/Medvedevsky/simple-web-application/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *APIServer) Start() error {

	if err := a.configureLogger(); err != nil {
		return err
	}

	a.configureRouter()

	if err := a.configureStore(); err != nil {
		return err
	}

	a.logger.Info("starting api server")

	// ListenAndServe - запускает http сервер с заданным адресом и обработчиком
	return http.ListenAndServe(a.config.WebAddress, a.router)
}

func (a *APIServer) configureLogger() error {
	// получаем уровень логирования
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)
	return nil
}

/*
		configureRouter - роутер, который будет обрабатывать входящий запрос
	 	и перенапровлять их к нужным обработчика
		в качестве роутер здесь gorilla/mux
*/
func (a *APIServer) configureRouter() {
	// HandleFunc - функция которая обрабатывает входящий запрос по эндпоинту
	a.router.HandleFunc("/hello", a.handleHello())
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ты чее тут забыл?")
	})
}

func (a *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (a *APIServer) configureStore() error {

	st := store.NewDb(a.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	a.store = st

	return nil
}
