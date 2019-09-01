package store_test

import (
	"github.com/go-redis/redis"
	"github.com/marko-ciric/tokens/store"
	"github.com/marko-ciric/tokens/util"
	models "gopkg.in/oauth2.v3/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	var s *store.RedisTokenStore
	BeforeEach(func() {
		redisClient := redis.NewClient(util.NewRedisOptions())
		s = store.NewTokenStore(redisClient)
	})
	Context("When token provided", func() {
		var token *models.Token
		BeforeEach(func() {
			token = models.NewToken()
			token.ClientID = "123"
		})
		It("Saves successfully", func() {
			Expect(s.Create(token)).To(BeNil())
		})
		It("Gets successfully", func() {
		})
	})

})
