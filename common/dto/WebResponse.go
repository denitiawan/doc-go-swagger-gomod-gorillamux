package dto

type WebResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data,omitempty"`
}

func NewWebResponse(code int, message string, errorMessage string, data interface{}) *WebResponse {
	response := &WebResponse{
		Code:         code,
		Message:      message,
		ErrorMessage: errorMessage,
		Data:         data,
	}

	return response
}
