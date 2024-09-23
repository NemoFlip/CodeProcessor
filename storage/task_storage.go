package storage

import (
	"HomeWork1/entity"
	"errors"
)

type TaskStorage struct {
	data map[string]entity.Task
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{data: make(map[string]entity.Task)}
}

func (rs *TaskStorage) Get(key string) (*entity.Task, error) {
	if val, ok := rs.data[key]; !ok {
		return nil, errors.New("there is no such key")
	} else {
		return &val, nil
	}
}

func (rs *TaskStorage) Put(key string, val entity.Task) error {
	rs.data[key] = val
	return nil
}

func (rs *TaskStorage) Post(key string, value entity.Task) error {
	if _, exists := rs.data[key]; exists {
		return errors.New("key is already exists")
	}
	rs.data[key] = value
	return nil
}

func (rs *TaskStorage) Delete(key string) error {
	if _, exists := rs.data[key]; !exists {
		return errors.New("key is not found")
	}
	delete(rs.data, key)
	return nil
}
