package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *API) Protect() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		token = c.GetHeader("Auth")
		if token == "" {
			var err error
			token, err = c.Cookie("Auth")
			if err != nil {
				c.JSON(401, gin.H{"error": "Auth token not found"})
				return
			}
		}
		user, err := a.auth.DecodeToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid Auth Token"})
			return
		}
		c.Set("userID", user.ID)
		c.Header("userID", strconv.FormatUint(uint64(user.ID), 10))
		c.Next()
	}
}
