package handlers

import (
	_ "HomeWork1/docs"
	"HomeWork1/entity"
	"HomeWork1/storage"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserServer struct {
	storage storage.UserStorage
}

func NewUserServer(storage storage.UserStorage) *UserServer {
	return &UserServer{storage: storage}
}

// @Summary Register User
// @Tags User
// @Description register new user
// @Accept json
// @Param user body entity.User true "Данные для регистрации пользователя"
// @Success 201
// @Failure 400
// @Router /register [post]
func (us *UserServer) RegisterHandler(ctx *gin.Context) {
	var newUser entity.User
	newUser.ID = uuid.New().String()
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := us.storage.Post(newUser) // TODO: Хэшировать пароли
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusCreated)
}

// @Summary Login User
// @Tags User
// @Description login registered user
// @Accept json
// @Param user body entity.User true "Данные для авторизации пользователя"
// @Success 200
// @Failute 400
// @Router /login [post]
func (us *UserServer) LoginHandler(ctx *gin.Context) {
	var user entity.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userFromDB, err := us.storage.Get(user.Login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.Password != userFromDB.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect password",
		})
		return
	}
	tokenCredentials := user.Login + ":" + user.Password
	token := base64.StdEncoding.EncodeToString([]byte(tokenCredentials))
	ctx.Writer.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
