package result

type (
	// Resp 通用返回结构体
	Resp struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		Data  *any   `json:"data"`
		Total *int   `json:"total"`
	}
)

// Success 成功返回
func Success(data any) *Resp {
	return &Resp{
		Code: CodeSuccess,
		Msg:  MsgMap[CodeSuccess],
		Data: &data,
	}
}

// Fail 错误返回
func Fail(code int, msg string) *Resp {
	return &Resp{
		Code: code,
		Msg:  msg,
	}
}
