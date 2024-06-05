package entity

type TxRequestBody struct {
	Id      string        `json:"id"`
	JsonRpc string        `json:"jsonRpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
