package controller

import (
	"denitiawan/research-swagger-gomod-gorillamux/common/dto"
	"encoding/json"
	"net/http"
)

func createWebResponse(httpStatusCode int, message string, data interface{}, errorMessage string, response http.ResponseWriter, request *http.Request) {

	// response body
	responseBody := dto.WebResponse{
		Code:         httpStatusCode,
		Message:      message,
		Data:         data,
		ErrorMessage: errorMessage,
	}

	// response header
	response.Header().Add("Content-Type", "application/json")

	// write response
	json.NewEncoder(response).Encode(responseBody)

}

func OK(message string, data interface{}, errorMessage string, response http.ResponseWriter, request *http.Request) {
	createWebResponse(http.StatusOK, message, data, errorMessage, response, request)
}

func BadRequest(message string, data interface{}, errorMessage string, response http.ResponseWriter, request *http.Request) {
	createWebResponse(http.StatusBadRequest, message, data, errorMessage, response, request)
}

func Error(message string, data interface{}, errorMessage string, response http.ResponseWriter, request *http.Request) {
	createWebResponse(http.StatusInternalServerError, message, data, errorMessage, response, request)
}
