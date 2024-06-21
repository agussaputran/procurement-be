package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	INVALID_AUTH_TOKEN = "invalid authorization token headers."
)

type JwtCustomClaims struct {
	Username string
	Email    string
	ID       string
	Role     string
	jwt.StandardClaims
}

type User struct {
	ID       string
	Username string
	Email    string
	Role     string
}

// CreateJWT func will used to create the JWT while signing in and signing out
func CreateTokenJWT(data User, secret string, expiry int) (response string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &JwtCustomClaims{
		Email: data.Email,
		ID:    data.ID,
		Role:  data.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

// IsAuthorized... validate token
func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractToken(token string) (string, error) {
	reqToken := token
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return "", errors.New(INVALID_AUTH_TOKEN)
	}
	return splitToken[1], nil
}

func ExtractClaims(token string) (jwt.MapClaims, error) {
	tokenString, err := ExtractToken(token)
	if err != nil {
		return nil, err
	}

	hmacSecretString := os.Getenv("ACCESS_TOKEN_SECRET") // value of jwt secret
	hmaSecret := []byte(hmacSecretString)

	payload, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return hmaSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if err := payload.Claims.Valid(); err != nil {
		return nil, err
	}

	return payload.Claims.(jwt.MapClaims), nil

}

func GetFullnameByToken(token string) (fullname string, success bool) {
	payload, check := ExtractClaims(token)
	if check != nil {
		return "", false
	}
	if payload["Name"] == nil {
		return "", false
	}
	fullname = payload["Name"].(string)
	return fullname, true
}

func GetIDByToken(token string) (id string, success bool) {
	payload, check := ExtractClaims(token)
	if check != nil {
		return "", false
	}
	if payload["ID"] == nil {
		return "", false
	}
	id = payload["ID"].(string)
	return id, true
}

func GetEmailByToken(token string) (email string, success bool) {
	payload, check := ExtractClaims(token)
	if check != nil {
		return "", false
	}
	if payload["Email"] == nil {
		return "", false
	}
	email = payload["Email"].(string)
	return email, true
}

func GetRole(token string) (id string, success bool) {
	payload, check := ExtractClaims(token)
	if check != nil {
		return "", false
	}
	if payload["Role"] == nil {
		return "", false
	}
	id = payload["Role"].(string)
	return id, true
}
