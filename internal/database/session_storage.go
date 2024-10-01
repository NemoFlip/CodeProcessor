package database

import (
	"HomeWork1/internal/entity"
	"errors"
)

type SessionStorage struct {
	data map[string]string
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{data: make(map[string]string)}
}

func (ss *SessionStorage) Get(userID string) (*entity.Session, error) {
	if value, exists := ss.data[userID]; exists {
		session := entity.Session{
			UserID:    userID,
			SessionID: value,
		}
		return &session, nil
	}
	return nil, errors.New("there is no such user")
}

func (ss *SessionStorage) Post(session entity.Session) {
	ss.data[session.UserID] = session.SessionID
}

func (ss *SessionStorage) Put(session entity.Session) error {
	if _, exists := ss.data[session.UserID]; exists {
		ss.data[session.UserID] = session.SessionID
		return nil
	}
	return errors.New("there is no such user")
}

func (ss *SessionStorage) Delete(session entity.Session) error {
	if _, exists := ss.data[session.UserID]; exists {
		delete(ss.data, session.UserID)
		return nil
	}
	return errors.New("there is no such user")
}
