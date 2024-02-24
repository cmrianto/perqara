package helper

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

type TokenPayload struct {
	UserId   int64
	UserName string
	ClientId *int64
	Lifespan int
	Secret   string
}

type AuthenticationData struct {
	UserId          int64
	UserName        string
	ClientId        *int64
	ContentTypeName string
	Permissions     []string
}

func GenerateToken(in TokenPayload) (string, error) {
	var (
		clientId int64 = 0
	)
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = in.UserId
	claims["user_name"] = in.UserName
	if in.ClientId != nil && *in.ClientId > 0 {
		clientId = *in.ClientId
		claims["client_id"] = clientId
	}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(in.Lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(in.Secret))

}

func TokenValid(c *gin.Context, in TokenPayload) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, Error(codes.Unauthenticated, "Unexpected signing method: ", nil)
		}
		return []byte(in.Secret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GetToken(c *gin.Context, in TokenPayload) (AuthenticationData, error) {
	result := AuthenticationData{}
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return result, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(in.Secret), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	result.UserId = int64(claims["user_id"].(float64))
	result.UserName = claims["user_name"].(string)
	if client_id, ok := claims["client_id"].(float64); ok && client_id > 0 {
		clientId := int64(client_id)
		result.ClientId = &clientId
	}
	return result, err
}
