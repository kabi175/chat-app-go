package domain

import "github.com/gin-gonic/gin"

type UserController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}
