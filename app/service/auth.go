package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"omochi/app/models"
	"strings"
	"time"
)

type TokenService struct {}

type Auth struct {
	UserId int64
	UserType int64
	Iat int64
}

const (
	secret = "VgHIOzK076FnCHA3NYgrJ2fZdfr9y5RRV5XBgwqvgNzNNop/7jC7Bg=="
	userIDKey = "user_id"
	userType = "type"
	iatKey =    "iat"
	expKey =    "exp"
	lifetime =  30 * time.Minute
)

func (TokenService) Generate(user *models.User, now time.Time) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDKey: user.ID,
		userType: user.Type,
		iatKey: now.Unix(),
		expKey: now.Add(lifetime).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func (TokenService) Verify(c *gin.Context) (*Auth, error){
	authHeader := c.Request.Header["Authorization"][0]
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2{
		authToken := bearerToken[1]
		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("token verify error")
			}
			return []byte(secret), nil
		})
		if err != nil {
			log.Printf("action=Verify err=%s", err.Error())
			return nil, err
		}
		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Printf("action=Verify not found claims in %s", authToken)
				return nil, fmt.Errorf("not found claims in %s", authToken)
			}
			userId, ok := claims[userIDKey].(float64)
			if !ok {
				log.Printf("action=Verify not found %s in %s", userIDKey, authToken)
				return nil, fmt.Errorf("not found %s in %s", userIDKey, authToken)
			}
			userType, ok := claims[userType].(float64)
			if !ok {
				log.Printf("action=Verify not found %s in %s", userType, authToken)
				return nil, fmt.Errorf("not found %s in %s", userType, authToken)
			}
			iat, ok := claims[iatKey].(float64)
			if !ok {
				log.Printf("action=Verify not found %s in %s", iatKey, authToken)
				return nil, fmt.Errorf("not found %s in %s", iatKey, authToken)
			}
			return &Auth{
				UserId: int64(userId),
				UserType: int64(userType),
				Iat: int64(iat),
			}, nil
		} else {
			log.Println("action=Verify failed to token valid")
			return nil, fmt.Errorf("failed to token valid")
		}
	} else {
		log.Println("action=Verify invalid token")
		return nil, fmt.Errorf("invalid token")
	}
}