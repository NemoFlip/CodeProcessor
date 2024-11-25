package handlers

import (
	_ "HomeWork1/docs"
	"HomeWork1/internal/database"
	"HomeWork1/internal/entity"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type UserServer struct {
	userStorage    database.UserStorage
	sessionStorage database.SessionStorage
}

func NewUserServer(userStorage database.UserStorage, sessionStorage database.SessionStorage) *UserServer {
	return &UserServer{userStorage: userStorage,
		sessionStorage: sessionStorage,
	}
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
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Printf("registration error: %s ", err)
		return
	}
	newUser.ID = uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 4)
	if err != nil {
		log.Printf("unable to hash the password")
		return
	}
	newUser.Password = string(hashedPassword)
	err = us.userStorage.Post(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err)
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
	userFromDB, err := us.userStorage.Get(user.Login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))
	if err != nil {
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
	newSession := entity.Session{
		UserID:    userFromDB.ID,
		SessionID: token,
	}
	err = us.sessionStorage.Post(newSession)
	if err != nil {
		log.Println(err)
	}

}
