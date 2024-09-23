package storage

import (
	"HomeWork1/entity"
	"errors"
)

type UserStorage struct {
	data []entity.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{data: make([]entity.User, 0)}
}

func (us *UserStorage) Post(newUser entity.User) error {
	for i := range us.data {
		if us.data[i].Login == newUser.Login {
			return errors.New("this user already exists")
		}
	}
	us.data = append(us.data, newUser)
	return nil
}

func (us *UserStorage) Get(username string) (*entity.User, error) {
	for i := range us.data {
		if username == us.data[i].Login {
			return &us.data[i], nil
		}
	}
	return nil, errors.New("this user doesn't exist")
}
