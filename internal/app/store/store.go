package store

import "database/sql"

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabseURL)
	if err != nil {
		return nil
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Cloce() {
	s.db.Close()
}
