package constant

type GlobolErrCode struct {
	Code    string `json:"code"`    // 异常编号
	Message string `json:"message"` // 异常信息 返回给前端显示
}
