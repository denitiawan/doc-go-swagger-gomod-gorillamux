package dto

import "denitiawan/research-swagger-gomod-gorillamux/common/logger"

type ImplResponse struct {
	Class        string
	Function     string
	Message      string
	Error        error
	ErrorMessage string
	Data         interface{}
}

func NewImplResponse(class string, function string, message string, error error, data interface{}) *ImplResponse {
	errorMessage := ""
	if error != nil {
		errorMessage = error.Error()
		logger.DoLogError(class, function, message, error)
	}

	response := &ImplResponse{
		Class:        class,
		Function:     function,
		Message:      message,
		Error:        error,
		ErrorMessage: errorMessage,
		Data:         data,
	}

	return response

}
