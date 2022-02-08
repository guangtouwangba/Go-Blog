package response

import "net/http"

type Response struct {
	Code    int64       `json:"code" default:"0"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) Success() *Response {
	r.Code = http.StatusOK
	return r
}

func (r *Response) SuccessWithData(data interface{}) *Response {
	r.Code = http.StatusOK
	r.Data = data
	return r
}

func (r *Response) Error(code int64, message string) *Response {
	r.Code = code
	r.Message = message
	return r
}
