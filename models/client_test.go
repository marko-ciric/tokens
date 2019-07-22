package models_test

import (
	"github.com/marko-ciric/tokens/models"
	"testing"
)

func TestClientStore(t *testing.T) {
	cli := models.Client{
		ID:     "",
		Secret: "",
		Domain: "",
		UserID: "",
	}
	s := ""
	models.Marshall(cli, &s)
	t.Log(s)
}

func TestRealClientStore(t *testing.T) {
	cli := models.Client{
		ID:     "123",
		Secret: "secret",
		Domain: "domain.com",
		UserID: "145",
	}
	s := ""
	models.Marshall(cli, &s)
	t.Log(s)
}
