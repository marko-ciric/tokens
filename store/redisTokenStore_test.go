package store_test

import (
	store "github.com/marko-ciric/tokens/store"
	"testing"
)

func TestRedisTokenStore(t *testing.T) {
	tokenStore := store.NewTokenStore()
	t.Log(tokenStore.GetByCode("123"))
}
