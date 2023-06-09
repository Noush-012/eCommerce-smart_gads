package response

import "strings"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ErrorResponse(statusCode int, message string, err string, data interface{}) Response {

	spiltedError := strings.Split(err, "\n")
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Errors:     spiltedError,
		Data:       data,
	}
}

func SuccessResponse(statusCode int, message string, data ...interface{}) Response {

	return Response{
		StatusCode: statusCode,
		Message:    message,
		Errors:     nil,
		Data:       data,
	}
}
