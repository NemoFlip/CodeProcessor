package database

import (
	"HomeWork1/internal/entity"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type SessionStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewSessionStorage() *SessionStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "1234",
		DB:       0,
	})
	return &SessionStorage{client: client, ctx: context.Background()}
}
func (ss *SessionStorage) Get(userID string) (string, error) {
	val, err := ss.client.Get(ss.ctx, userID).Result()
	if err != nil {
		return "", fmt.Errorf("unable to get value by userID(%s): %w", userID, err)
	}
	return val, nil
}

func (ss *SessionStorage) Post(session entity.Session) error {
	err := ss.client.Set(ss.ctx, session.UserID, session.SessionID, 0).Err()
	if err != nil {
		return fmt.Errorf("unable to set value: %w", err)
	}
	return nil
}

//func (ss *SessionStorage) Put(session entity.Session) error {
//	if _, exists := ss.data[session.UserID]; exists {
//		ss.data[session.UserID] = session.SessionID
//		return nil
//	}
//	return errors.New("there is no such user")
//}
//
//func (ss *SessionStorage) Delete(session entity.Session) error {
//	if _, exists := ss.data[session.UserID]; exists {
//		delete(ss.data, session.UserID)
//		return nil
//	}
//	return errors.New("there is no such user")
//}
