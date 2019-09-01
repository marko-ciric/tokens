package store

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"gopkg.in/oauth2.v3"
)

// NewTokenStore create client store
func NewTokenStore(client *redis.Client) *RedisTokenStore {
	return &RedisTokenStore{client: client}
}

// RedisTokenStore client information store
type RedisTokenStore struct {
	sync.RWMutex
	client *redis.Client
}

func (store *RedisTokenStore) Create(info oauth2.TokenInfo) error {
	store.Lock()
	defer store.Unlock()

	return store.client.Set(info.GetCode(), info.GetAccess(), 0).Err()
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
