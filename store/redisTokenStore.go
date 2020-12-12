package store

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"gopkg.in/oauth2.v3"
	models "gopkg.in/oauth2.v3/models"
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

//Create and persist an instance of TokenInfo
func (store *RedisTokenStore) Create(info oauth2.TokenInfo) error {
	fmt.Printf("Token received %s", info)
	store.Lock()
	defer store.Unlock()
	byteArray, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("Error encoding token for client %s", info.GetClientID())
		return err
	}

	err = store.client.Set(info.GetClientID(), byteArray, time.Hour).Err()
	if err != nil {
		fmt.Printf("Error setting clientId token mapping for %s", info.GetClientID())
		return err
	}
	err = store.client.Set(info.GetAccess(), byteArray, time.Hour).Err()
	if err != nil {
		fmt.Printf("Error setting access token mapping for %s", info.GetClientID())
		return err
	}
	err = store.client.Set(info.GetCode(), byteArray, time.Hour).Err()
	if err != nil {
		fmt.Printf("Error setting authorisation code token mapping for %s", info.GetAccess())
		return err
	}
	err = store.client.Set(info.GetRefresh(), byteArray, time.Hour).Err()
	if err != nil {
		fmt.Printf("Error setting refresh token mapping for %s", info.GetCode())
		return err
	}
	return nil
}

//RemoveByCode finds a token by code and removes it from Redis
func (store *RedisTokenStore) RemoveByCode(code string) error {
	return store.client.Del(code).Err()
}

//RemoveByAccess finds a token by code and removes it from Redis
func (store *RedisTokenStore) RemoveByAccess(access string) error {
	return store.client.Del(access).Err()
}

//RemoveByRefresh finds a token by code and removes it from Redis
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
	var info models.Token
	err = json.Unmarshal([]byte(val), &info)
	return &info, err
}

func (store *RedisTokenStore) getClientInfoByKey(key string) (oauth2.ClientInfo, error) {
	var val string
	val, err := store.getStringByKey(key)
	if err != nil {
		panic(err)
	}
	var info models.Client
	err = json.Unmarshal([]byte(val), &info)
	return &info, err
}

// GetByCode returns a valid access token by a given authorization_code
func (store *RedisTokenStore) GetByCode(code string) (oauth2.TokenInfo, error) {
	return store.getTokenInfoByKey(code)
}

// GetByAccess returns a valid access token by a given authorization_code
func (store *RedisTokenStore) GetByAccess(access string) (oauth2.TokenInfo, error) {
	return store.getTokenInfoByKey(access)
}

// GetByRefresh returns a valid access token by a given authorization_code
func (store *RedisTokenStore) GetByRefresh(refresh string) (oauth2.TokenInfo, error) {
	return store.getTokenInfoByKey(refresh)
}
