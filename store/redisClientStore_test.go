package store_test

import (
	"github.com/marko-ciric/tokens/store"
	"testing"
)

func TestRedisClientStore(t *testing.T) {
	clientStore := store.NewClientStore()
	t.Log(clientStore.GetByID("1123"))
}
