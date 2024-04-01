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
	TelegramHandle string `json:"telegram_handle"`
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

func (service *JwtService) TokenGenerate(username string, roleID int, telegramHandle string) (string, error) {
	expirationTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		log.Fatalln(err)
	}
	claims := &authCustomClaims{
		Username: username,
		RoleID:   roleID,
		TelegramHandle: telegramHandle,
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
		fmt.Println("Token validated:", claims.Username, claims.RoleID, claims.Issuer, claims.TelegramHandle)
		return nil
	} else {
		return fmt.Errorf("invalid token or claims")
	}
}

func (service *JwtService) JwtAuthMiddlewareDefault() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.TokenValidate(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorised, please login"})
			c.Abort()
			return
		}
		c.Next()
	}
}


func DecodeToken(c *gin.Context) (*authCustomClaims, error) {
	tokenString := ExtractToken(c)
	privateKey := os.Getenv("PRIVATE_KEY") // Ensure the private key is available
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(privateKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token or claims")
	}
}


func CheckAdmin(claims *authCustomClaims) error {
	if claims.RoleID != 2 {
			return fmt.Errorf("only admin can access")
		}
	return nil
}

func CheckMatchingCallerUsername(c *gin.Context, username string) error {
	claims ,err := DecodeToken(c)
	if err != nil {
		log.Println("failed to decode auth token on header")
		return fmt.Errorf("failed to decode auth token for checking caller id")
	}
	// admin auto access
	if err := CheckAdmin(claims); err != nil {
		return err
	}
	if claims.Username != username {
		return fmt.Errorf("unauthorised user access")
	}
	return nil
}

func CheckMatchingCallerId(c *gin.Context, userId string) error {
	claims ,err := DecodeToken(c)
	if err != nil {
		log.Println("failed to decode auth token on header")
		return fmt.Errorf("failed to decode auth token for checking caller id")
	}
	// admin auto access
	if err := CheckAdmin(claims); err != nil {
		return err
	}
	if claims.ID != userId {
		return fmt.Errorf("unauthorised user access")
	}
	return nil
}

func (service *JwtService) JwtAuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.TokenValidate(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorised, please login"})
			c.Abort()
			return
		}
		// check admin
		claims ,err := DecodeToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		if err := CheckAdmin(claims); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}