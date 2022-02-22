package res

import (
	"net/http"
)

type ResponseModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}

func Response(data interface{}, message string) ResponseModel {
	result := ResponseModel{
		Success: true, Data: data, Message: message, Code: http.StatusOK,
	}

	return result
}

func ResponseError(message string, code int) ResponseModel {
	result := ResponseModel{
		Success: false, Data: nil, Message: message, Code: code,
	}

	return result
}
