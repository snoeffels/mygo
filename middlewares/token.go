package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		mode := os.Getenv("GIN_MODE")
		if mode == "release" {

			token := c.Request.Header.Get("Authorization")
			if token == "" {
				c.JSON(http.StatusForbidden, gin.H{"error": "No Authorization header provided"})
				c.Abort()
				return
			}

			_, err := tokenValid(c.Request)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}

			//roles := claims["roles"].([]interface{})
			//strRoles := make([]string, len(roles))
			//for i, v := range roles {
			//	strRoles[i] = fmt.Sprint(v)
			//}
			//c.Set("roles", strRoles)

			c.Next()
		}

		c.Next()
	}
}

func tokenValid(r *http.Request) (jwt.MapClaims, error) {
	token, err := verifyJWTToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, err
	}

	return claims, nil
}

func verifyJWTToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	jwtAccSecret := os.Getenv("JWT_ACCESS_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtAccSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	strArr := strings.Split(token, " ")

	if len(strArr) == 2 {
		if strArr[0] == "Bearer" {
			return strArr[1]
		}
	}

	return ""
}
