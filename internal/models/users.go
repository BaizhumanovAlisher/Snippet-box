package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	//todo
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	//todo
	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	//todo
	return false, nil
}
