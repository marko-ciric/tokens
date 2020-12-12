package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"

	"github.com/marko-ciric/tokens/store"
	"github.com/marko-ciric/tokens/util"
)

func main() {
	redisClient := redis.NewClient(util.NewRedisOptions())
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewTokenStore(redisClient), nil)

	// client memory store
	clientStore := store.NewClientStore(redisClient)
	err := clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	if err != nil {
		panic(err)
	}
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)

	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}
