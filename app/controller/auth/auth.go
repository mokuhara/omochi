package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"omochi/app/models"
	"omochi/app/repository"
	"omochi/app/service"
	"time"
)

func Signup(c *gin.Context){
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	purePassword := user.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("failed to create password hash")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)
	userRepository := repository.UserRepository{}
	err = userRepository.Create(&user)
	if err != nil {
		log.Println("failed to create user")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	//passwordをhash化前に戻す
	user.Password = purePassword
	token, err :=  createToken(&user, c)
	if err != nil {
		log.Println("failed to create token")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}

	data := map[string]string{"token": token}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": data,
	})
}

func Login(c *gin.Context){
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	token, err := createToken(&user, c)
	if err != nil {
		log.Println("failed to create token")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}

	data := map[string]string{"token": token}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": data,
	})
}


func createToken(user *models.User, c *gin.Context) (string, error){
	userRepository := repository.UserRepository{}
	matchUser, err := userRepository.GetByEmail(user.Email)
	if err != nil {
		log.Println("failed to get user by email")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(matchUser.Password), []byte(user.Password))
	if err != nil {
		log.Println("invalid password")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return "", err
	}

	tokenService := service.TokenService{}
	token, err := tokenService.Generate(matchUser, time.Now())
	if err != nil {
		log.Println("failed create token")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return "", err
	}
	return token, nil
}