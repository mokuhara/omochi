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
		log.Println("action=Signup bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}

	purePassword := user.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("action=Signup failed to create password hash")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	user.Password = string(hash)

	switch user.Type {
		case 1:
			specialistRepository := repository.SpecialistRepository{}
			err = specialistRepository.Create(&user)
			if err != nil {
				log.Println("action=Signup failed to create specialist")
				c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
				return
			}
		case 2:
			clientRepository := repository.ClientRepository{}
			err = clientRepository.Create(&user)
			if err != nil {
				log.Println("action=Signup failed to create client")
				c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
				return
			}
	}
	//passwordをhash化前に戻す
	user.Password = purePassword
	resAuth, err :=  createToken(&user, c)
	if err != nil {
		log.Println("action=Signup failed to create token")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{"userId": resAuth.userId, "token":resAuth.token, "type": resAuth.userType}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": data,
	})
}

func Login(c *gin.Context){
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		log.Println("action=Login bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	resAuth, err := createToken(&user, c)
	if err != nil {
		log.Println("action=Login failed to create token")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{"userId": resAuth.userId, "token":resAuth.token, "type": resAuth.userType}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": data,
	})
}

type resAuth struct {
	userId    int64
	userType  models.Type
	token     string
}


func createToken(user *models.User, c *gin.Context) (*resAuth, error){
	userRepository := repository.UserRepository{}
	matchUser, err := userRepository.GetByEmail(user.Email)
	if err != nil {
		log.Println("action=createToken failed to get user by email")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(matchUser.Password), []byte(user.Password))
	if err != nil {
		log.Println("action=createToken invalid password")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return nil, err
	}

	tokenService := service.TokenService{}
	token, err := tokenService.Generate(matchUser, time.Now())
	if err != nil {
		log.Println("action=createToken failed create token")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return nil, err
	}

	resAuth := resAuth{token: token, userId: int64(matchUser.ID), userType: matchUser.Type}
	return &resAuth, nil
}

func Router(group *gin.RouterGroup){
	authEngine := group.Group("/auth")
	{
		authEngine.POST("/signup", Signup)
		//authEngine.Use(middleware.IsExistsUserInfo())
		authEngine.POST("/login", Login)
	}
}