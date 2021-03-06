package store_test

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/marko-ciric/tokens/store"
	"github.com/marko-ciric/tokens/util"
	"gopkg.in/oauth2.v3"
	models "gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/utils/uuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func convertJSON(token oauth2.TokenInfo) string {
	b, _ := json.Marshal(token)
	return string(b)
}
func compareJSON(expected oauth2.TokenInfo, actual oauth2.TokenInfo) {
	Expect(convertJSON(expected)).To(MatchJSON(convertJSON(actual)))
}

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
		clientId, _ := uuid.NewRandom()
		token.ClientID = clientId.String()
		token.Code = "123"
		token.Scope = "read"
		err = s.Create(token)
		Expect(err).NotTo(HaveOccurred())
	})
	Context("gets by access code", func() {
		BeforeEach(func() {
			persistedToken, err = s.GetByCode(token.Code)
		})
		It("returns valid token", func() {
			compareJSON(token, persistedToken)
		})
	})
	Context("gets by access token", func() {
		BeforeEach(func() {
			persistedToken, err = s.GetByAccess(token.Access)
		})
		It("returns valid token", func() {
			compareJSON(token, persistedToken)
		})
	})
	Context("gets by refresh token", func() {
		BeforeEach(func() {
			persistedToken, err = s.GetByRefresh(token.Refresh)
		})
		It("returns valid token", func() {
			compareJSON(token, persistedToken)
		})
	})
})
