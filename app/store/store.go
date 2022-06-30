package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	config   *Config
	db       sql.DB
	userRepo *UserRepo
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s",
		s.config.DbUser, s.config.DbPassword, s.config.DbName, s.config.DbPort, s.config.SslMode)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = *db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepo {
	if s.userRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepo{
		store: s,
	}

	return s.userRepo
}
