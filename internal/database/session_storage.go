package database

import (
	"HomeWork1/configs"
	"HomeWork1/internal/entity"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type SessionStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewSessionStorage(cfg configs.Config) *SessionStorage {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.ServerMain.DatabaseRedis.Host,
			cfg.ServerMain.DatabaseRedis.Port),
		Password: cfg.ServerMain.DatabaseRedis.Password,
		DB:       cfg.ServerMain.DatabaseRedis.DB,
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
		return fmt.Errorf("unable to insert value by userID(%s): %w", session.UserID, err)
	}
	return nil
}

func (ss *SessionStorage) Put(session entity.Session) error {
	err := ss.client.Set(ss.ctx, session.UserID, session.SessionID, 0).Err()
	if err != nil {
		return fmt.Errorf("unable to update value by userID(%s): %w", session.UserID, err)
	}
	return nil
}

func (ss *SessionStorage) Delete(userID string) error {
	result, err := ss.client.Del(ss.ctx, userID).Result()
	if err != nil {
		return fmt.Errorf("unable to delete by userID(%s): %w", userID, err)
	}
	if result == 0 {
		return fmt.Errorf("unable to delete: there is no such key(%s)", userID)
	}
	return nil
}
