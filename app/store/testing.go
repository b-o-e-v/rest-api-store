package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(
	t *testing.T,
	DbHost string,
	DbPort string,
	DbName string,
	DbUser string,
	DbPassword string,
	SslMode string,
) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()

	config.DbHost = DbHost
	config.DbPort = DbPort
	config.DbName = DbName
	config.DbUser = DbUser
	config.DbPassword = DbPassword
	config.SslMode = SslMode

	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
