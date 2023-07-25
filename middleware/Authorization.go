package middleware

import (
	baseController "denitiawan/research-swagger-gomod-gorillamux/common/controller"
	"denitiawan/research-swagger-gomod-gorillamux/config"
	"denitiawan/research-swagger-gomod-gorillamux/security"
	"log"
	"net/http"
	"strings"
)

func Authorization(appConfig config.AppConfig, next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var token string

		// get token from http header
		authorizationHeader := request.Header.Get("Authorization")

		// replace prefix on jwt token
		token = strings.Replace(authorizationHeader, "Bearer ", "", 1)

		// validate : token is empty
		if token == "" {
			baseController.Error("Token is Empty", nil, "", response, request)
			return
		}

		// validate : expired time, jwt format, get token value, etc
		jwtValue, err := security.ValidateToken(token, appConfig.TokenSecret)
		if err != nil {
			baseController.Error("Token is Invalid", nil, err.Error(), response, request)
			return
		}

		// get current user
		userId, _ := jwtValue.(float64)
		log.Printf("User %s logged in.", userId)

		// set user into ctx
		// todo

		// redirect to controller
		next(response, request)

	}
}
