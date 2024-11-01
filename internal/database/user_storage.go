package database

import (
	"HomeWork1/internal/entity"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type UserStorage struct {
	DB *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{DB: db}
}

func (us *UserStorage) Post(newUser entity.User) error {
	if err := us.DB.Ping(); err != nil {
		return fmt.Errorf("database connection is closed: %w", err)
	}

	query := "INSERT INTO users (id, login, password) VALUES($1, $2, $3)"
	result, err := us.DB.Exec(query, newUser.ID, newUser.Login, newUser.Password)
	if err != nil {
		return fmt.Errorf("unable to insert new user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were inserted")
	}
	return nil
}

func (us *UserStorage) Get(name string) (*entity.User, error) {
	if err := us.DB.Ping(); err != nil {
		fmt.Printf("database connection is closed in get: %s", err)
	}
	query := "SELECT id, login, password FROM users WHERE login = $1"

	row := us.DB.QueryRow(query, name)

	var user entity.User

	err := row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with name(%s) has not found", name)
		}
		return nil, fmt.Errorf("unable to scan user: %w", err)
	}
	return &user, nil
}
