package middleware

import (
	"errors"
	"github.com/adityarizkyramadhan/tes_intern_delos/infrastructure/app"
	"github.com/adityarizkyramadhan/tes_intern_delos/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func GenerateToken(id uint) (string, error) {
	envApp, err := app.NewDriverApp()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		//"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(envApp.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ValidateJWToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail(
				"Authorization is empty",
				errors.New("bearer token is not found")))
			return
		}
		cleanToken := strings.ReplaceAll(bearerToken, "Bearer ", "")
		token, err := jwt.Parse(cleanToken, ekstractToken)
		if err != nil {
			c.JSON(http.StatusForbidden, response.ResponseWhenFail(
				"Middleware authentication failed when parsing",
				err.Error()))
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := uint(claims["id"].(float64))
			c.Set("login", userId)
		} else {
			c.JSON(http.StatusForbidden, response.ResponseWhenFail(
				"Middleware authentication failed when claims",
				err))
			return
		}
	}
}

func ekstractToken(token *jwt.Token) (interface{}, error) {
	envApp, err := app.NewDriverApp()
	if err != nil {
		return "", err
	}
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(envApp.SecretKey), nil
}
