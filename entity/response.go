package entity

const (
	OK                  = 0
	JsonDecodeErr       = 100001
	NotSupportTransType = 100002
	InitSyncErr         = 100003
	GetTransactionErr   = 100004
)

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func GetErrResponse(code int, msg string) (resp *Response) {
	if msg == "" {
		resp = &Response{Code: code, Msg: "", Result: ""}
	} else {
		resp = &Response{Code: code, Msg: "", Result: ""}
	}
	return resp
}

func GetResponse(code int, msg string, data interface{}) (resp *Response) {
	if msg == "" {
		resp = &Response{Code: code, Msg: "", Result: data}
	} else {
		resp = &Response{Code: code, Msg: msg, Result: data}
	}
	return resp
}
