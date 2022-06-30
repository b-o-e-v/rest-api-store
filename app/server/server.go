package server

import (
	"io"
	"net/http"

	"github.com/b-o-e-v/rest-api-store/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// СТРУКТУРА СЕРВЕРА
type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// СОЗДАЕМ НОВУЮ КОНФИГУРАЦИЮ СЕРВЕРА
func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// ЗАПУСКАЕМ СЕРВЕР
func (s *Server) Start() error {
	if err := s.configureLoger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Сервер запущен!")
	s.logger.Info("Postgres started at PORT: ", s.config.Store.DbPort)

	return http.ListenAndServe(s.config.Port, s.router)
}

// КОНФИГУРИРУЕМ ЛОГГЕР
func (s *Server) configureLoger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// КОНФИГУРИРУЕМ РОУТЕР
func (s *Server) configureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *Server) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *Server) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
