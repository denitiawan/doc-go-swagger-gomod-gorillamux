package controller

import (
	"denitiawan/research-swagger-gomod-gorillamux/common/dto"
	"encoding/json"
	"net/http"
)

func CreateWebResponse(httpStatusCode int, status string, data interface{}, response http.ResponseWriter, request *http.Request) {

	// response body
	responseBody := dto.WebResponse{
		Code:   httpStatusCode,
		Status: status,
		Data:   data,
	}

	// response header
	response.Header().Add("Content-Type", "application/json")

	// write response
	json.NewEncoder(response).Encode(responseBody)

}
