package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/kabi175/chat-app-go/messager/domain/apperrors"
)

type GinUserController struct {
	userService  domain.UserService
	tokenService domain.TokenService
}

func NewGinUserController(userService domain.UserService, tokenService domain.TokenService) domain.UserController {
	return &GinUserController{
		userService:  userService,
		tokenService: tokenService,
	}
}

func (g *GinUserController) Post(c *gin.Context) {
	var user *domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := g.userService.SignUp(user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (g *GinUserController) Get(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := g.userService.GetByID(uint(userID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (g *GinUserController) Put(c *gin.Context) {
	var user *domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := g.userService.LogIn(user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("Auth", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (g *GinUserController) Delete(c *gin.Context) {
	userIDValue, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.NewInternalServerError("userID not found")})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.NewInternalServerError("userID is not of type uint")})
		return
	}
	user := domain.User{ID: uint(userID)}
	err := g.userService.Delete(&user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
