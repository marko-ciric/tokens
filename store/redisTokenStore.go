package store

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis"
	"gopkg.in/oauth2.v3"
)

// NewTokenStore create client store
func NewTokenStore() *RedisTokenStore {
	return &RedisTokenStore{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

// RedisTokenStore client information store
type RedisTokenStore struct {
	sync.RWMutex
	client *redis.Client
}

func (store *RedisTokenStore) Create(info oauth2.TokenInfo) error {
	store.Lock()
	defer store.Unlock()

	return store.client.Set(info.GetCode(), info, 0).Err()
}

func (store *RedisTokenStore) RemoveByCode(code string) error {
	return store.client.Del(code).Err()
}

func (store *RedisTokenStore) RemoveByAccess(access string) error {
	return store.client.Del(access).Err()
}

func (store *RedisTokenStore) RemoveByRefresh(refresh string) error {
	return store.client.Del(refresh).Err()
}

func (store *RedisTokenStore) getStringByKey(key string) (string, error) {
	store.RLock()
	defer store.RUnlock()
	var val string
	val, err := store.client.Get(key).Result()
	if err == redis.Nil {
		fmt.Printf("Token %s does not exist", key)
		return "", err
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (store *RedisTokenStore) getTokenInfoByKey(key string) (oauth2.TokenInfo, error) {
	var val string
	val, err := store.getStringByKey(key)
	if err != nil {
		panic(err)
	}
	var info oauth2.TokenInfo
	json.Unmarshal([]byte(val), &info)
	return info, nil
}

func (store *RedisTokenStore) getClientInfoByKey(key string) (oauth2.ClientInfo, error) {
	var val string
	val, err := store.getStringByKey(key)
	if err != nil {
		panic(err)
	}
	var info oauth2.ClientInfo
	json.Unmarshal([]byte(val), &info)
	return info, nil
}

func (store *RedisTokenStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	return store.getTokenInfoByKey(code)
}

func (store *RedisTokenStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	return store.getTokenInfoByKey(access)
}

func (store *RedisTokenStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	return store.getTokenInfoByKey(refresh)
}

// GetByID according to the ID for the client information
func (store *RedisTokenStore) GetByID(id string) (cli oauth2.ClientInfo, err error) {
	return store.getClientInfoByKey(id)
}

// Set set client information
func (cs *RedisTokenStore) Set(id string, cli oauth2.ClientInfo) (err error) {
	cs.Lock()
	defer cs.Unlock()

	if err := cs.client.Set(id, cli, 0).Err(); err != nil {
		panic(err)
	}
	return err
}
