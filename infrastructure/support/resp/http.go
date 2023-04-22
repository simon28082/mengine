package resp

import "net/http"

func SuccessWithData(data interface{}) *Response {
	return &Response{
		Status: http.StatusOK,
		Data:   data,
	}
}

func Success() *Response {
	return &Response{
		Status: http.StatusOK,
	}
}

func Failed(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}

func Error(status int, err error) *Response {
	return Failed(status, err.Error())
}
