package response

type Response struct {
	Status  string      `json:"status" default:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) SetStatus(status string) {
	r.Status = status
}
func (r *Response) Success() *Response {
	r.Status = "ok"
	return r
}
