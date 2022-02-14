package middleware

import (
	"fmt"
	"log"
	"net/http"

	"majoo/response"
	"majoo/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println("token from header",authHeader)
		if authHeader == "" {
			response := response.BuildErrorResponse("Failed to process request", "No token provided", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader, c)
		if err != nil {
			response := response.BuildErrorResponse("Error", "Your token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		fmt.Println("ini token",token, token.Valid)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			fmt.Println("masuk sini")
			response := response.BuildErrorResponse("Error", "Your token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
