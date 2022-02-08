package response

type Response struct {
	Code    int64       `json:"code" default:"0"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) Success() *Response {
	r.Code = 0
	return r
}

func (r *Response) SuccessWithData(data interface{}) *Response {
	r.Code = 0
	r.Data = data
	return r
}
