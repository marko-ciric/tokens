package store_test

import (
	"github.com/go-redis/redis"
	"github.com/marko-ciric/tokens/store"
	"github.com/marko-ciric/tokens/util"
	"gopkg.in/oauth2.v3"
	models "gopkg.in/oauth2.v3/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("when token provided", func() {
	var (
		token          *models.Token
		persistedToken oauth2.TokenInfo
		s              *store.RedisTokenStore
		err            error
	)
	BeforeEach(func() {
		redisClient := redis.NewClient(util.NewRedisOptions())
		s = store.NewTokenStore(redisClient)
		token = models.NewToken()
		token.ClientID = "123"
		token.Code = "123"
		token.Scope = "read"
		err = s.Create(token)
		Expect(err).NotTo(HaveOccurred())
	})
	Context("gets by access code", func() {
		BeforeEach(func() {
			token = models.NewToken()
			token.ClientID = "124"
			token.Code = "123"
			token.Scope = "read"
			err = s.Create(token)
			Expect(err).NotTo(HaveOccurred())
			persistedToken, err = s.GetByCode(token.Code)
		})
		It("access code", func() {
			Expect(persistedToken.GetCode()).To(Equal(token.Code))
		})
		It("clientId", func() {
			Expect(persistedToken.GetClientID()).To(Equal(token.ClientID))
		})
		It("right scope", func() {
			Expect(persistedToken.GetScope()).To(Equal(token.Scope))
		})
	})
})
