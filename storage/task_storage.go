package storage

import (
	"HomeWork1/entity"
	"errors"
)

type RamStorage struct {
	data map[string]entity.Task
}

func NewRamStorage() *RamStorage {
	return &RamStorage{data: make(map[string]entity.Task)}
}

func (rs *RamStorage) Get(key string) (*entity.Task, error) {
	if val, ok := rs.data[key]; !ok {
		return nil, errors.New("there is no such key")
	} else {
		return &val, nil
	}
}

func (rs *RamStorage) Put(key string, val entity.Task) error {
	rs.data[key] = val
	return nil
}

func (rs *RamStorage) Post(key string, value entity.Task) error {
	if _, exists := rs.data[key]; exists {
		return errors.New("key is already exists")
	}
	rs.data[key] = value
	return nil
}

func (rs *RamStorage) Delete(key string) error {
	if _, exists := rs.data[key]; !exists {
		return errors.New("key is not found")
	}
	delete(rs.data, key)
	return nil
}
