package usecase

import (
	e "github.com/inagacky/go_gin_sample/src/api/usecase/error"
)

type CommonResponse struct {
	Status int `json:"status"`
	Error *e.ErrorResponse `json:"error"`
	Result interface{} `json:"result"`
}

// バリデートのエラーレスポンスを作成する
func (re *CommonResponse) CreateValidateErrorResponse(error string) *CommonResponse {
	re.Status = StatusError
	// ErrorResponseの定義
	errorResponse := &e.ErrorResponse{}
	errorResponse.Code = e.ErrorCodeRequestValidate
	errorResponse.Error = error
	re.Error = errorResponse

	return re
}

// SQLエラーのレスポンスを生成する
func (re *CommonResponse) CreateSQLErrorResponse(error string) *CommonResponse {
	re.Status = StatusError
	// ErrorResponseの定義
	errorResponse := &e.ErrorResponse{}
	errorResponse.Code = e.ErrorCodeSQL
	errorResponse.Error = error
	re.Error = errorResponse

	return re
}

// 正常時のレスポンスを作成する
func (re *CommonResponse) CreateSuccessResponse(result interface{}) *CommonResponse {
	re.Status = StatusSuccess
	re.Result = result
	re.Error = nil

	return re
}

// レスポンスステータス
const (
	StatusSuccess = 200 // 正常時
	StatusError = 500   // エラー時
)