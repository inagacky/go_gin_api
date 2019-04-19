package usecase

import (
	e "github.com/go_gin_sample/apps/usecase/error"
)

type CommonResponse struct {
	Status int `json:"status"`
	Error e.ErrorResponse `json:"error"`
	Result struct{} `json:"result"`
}

// バリデートのエラーレスポンスを作成する
func (re *CommonResponse) CreateValidateErrorResponse(error string) *CommonResponse {
	re.Status = StatusError
	// ErrorResponseの定義
	errorResponse := e.ErrorResponse{}
	errorResponse.Code = e.ErrorCodeRequestValidate
	errorResponse.Error = error
	re.Error = errorResponse

	return re
}

//
const (
	StatusSuccess = 200
	StatusError = 500
)