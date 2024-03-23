package auth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type authCustomClaims struct {
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
	jwt.RegisteredClaims
}

type JwtService struct {
	privateKey string
	issuedBy   string
}

func JWTAuthService() *JwtService {
	return &JwtService{
		privateKey: os.Getenv("PRIVATE_KEY"),
		issuedBy:   "jk",
	}
}

func (service *JwtService) TokenGenerate(username string, roleID int) (string, error) {
	expirationTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		log.Fatalln(err)
	}
	claims := &authCustomClaims{
		Username: username,
		RoleID:   roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expirationTime))),
			Issuer:    service.issuedBy,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(service.privateKey))

	return signedToken, err
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorisation")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
func (service *JwtService) TokenValidate(c *gin.Context) error {
	jwtTokenStr := ExtractToken(c)
	token, err := jwt.ParseWithClaims(jwtTokenStr, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(service.privateKey), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		fmt.Println("Token validated:", claims.Username, claims.RoleID, claims.Issuer)
		return nil
	} else {
		return fmt.Errorf("invalid token or claims")
	}
}

func (service *JwtService) JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.TokenValidate(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorised. Go be authorised den come back")
			c.Abort()
			return
		}
		c.Next()
	}
}
