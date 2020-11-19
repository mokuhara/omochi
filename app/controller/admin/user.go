package admin

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"omochi/app/models"
	"omochi/app/repository"
	"strconv"
)

func GetUsers(c *gin.Context){
	userRepository := repository.UserRepository{}
	users, err := userRepository.GetAll()
	if err != nil {
		log.Println("action=GetUsers failed to get users")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": users,
	})
}

func UpdateUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("action=UpdateUser bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	userRepository := repository.UserRepository{}
	err = userRepository.Update(&user)
	if err != nil {
		log.Println("action=UpdateUser failed to update user")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.PostForm("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("action=DeleteUser failed to parse user_id")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	userRepository := repository.UserRepository{}
	err = userRepository.Delete(intId)
	if err != nil {
		log.Println("action=DeleteUser failed to delete user")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("action=CreateUser bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("action=CreateUser failed to create hash")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)
	userRepository := repository.UserRepository{}
	err = userRepository.Create(&user)
	if err != nil {
		log.Println("action=CreateUser failed to create user")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": user,
	})
}
