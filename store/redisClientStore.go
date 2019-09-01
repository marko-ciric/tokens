package store

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"github.com/marko-ciric/tokens/models"
	"gopkg.in/oauth2.v3"
)

// NewClientStore create client store
func NewClientStore(client *redis.Client) *RedisClientStore {
	return &RedisClientStore{
		client: client,
	}
}

// ClientStore client information store
type RedisClientStore struct {
	sync.RWMutex
	client *redis.Client
}

// GetByID according to the ID for the client information
func (store *RedisClientStore) GetByID(id string) (cli oauth2.ClientInfo, err error) {
	store.RLock()
	defer store.RUnlock()
	c, err := store.client.Get(id).Result()
	if err == redis.Nil {
		fmt.Println("Key does not exist")
	}
	if err != nil {
		panic(err)
	}
	err = models.Unmarshall(&cli, c)
	if err != nil {
		panic(err)
	}
	return
}

// Set set client information
func (cs *RedisClientStore) Set(id string, cli oauth2.ClientInfo) (err error) {
	cs.Lock()
	defer cs.Unlock()
	if err := cs.client.Set(id, cli, 0).Err(); err != nil {
		panic(err)
	}
	return nil
}
