package store_test

import (
	"github.com/marko-ciric/tokens/store"
	models "gopkg.in/oauth2.v3/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	Context("When token provided", func() {
		var s *store.RedisTokenStore
		var token *models.Token
		BeforeEach(func() {
			s = store.NewTokenStore()
			token = models.NewToken()
			token.ClientID = "123"
		})
		It("Saves successfully", func() {
			Expect(s.Create(token)).To(Equal("done"))
		})
		It("Gets successfully", func() {
		})
	})

})
