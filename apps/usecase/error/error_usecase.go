package usecase


// 共通構造体の定義
type ErrorResponse struct {
	Code int `json:"code"`
	Error string `json:"error"`
}

const (
	ErrorCodeRequestValidate = 100
)