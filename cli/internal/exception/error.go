package exception

import "fmt"

type Error struct {
	errCode uint16
	message string
	param   map[string]string
}

// NewFromErrCode エラーコードからErrorを生成
// 存在しないエラーコードを指定した場合は FAILED_MAKE_ERROR のエラーを返す
func NewFromErrCode(errCode uint16) *Error {
	if !existsErrCode(errCode) {
		return NewFromErrCode(FAILED_MAKE_ERROR)
	}
	return newErr(errCode, errList[errCode])
}

// NewParamErrFromCode エラーコードとパラメータからErrorを生成
// 存在しないエラーコードを指定した場合は FAILED_MAKE_ERROR のエラーを返す
func NewParamErrFromCode(errCode uint16, param map[string]string) *Error {
	if !existsErrCode(errCode) {
		return NewFromErrCode(FAILED_MAKE_ERROR)
	}
	return newParamErr(errCode, errList[errCode], param)
}

// Error エラーメッセージを出力する
// 出力フォーマットは [code] <Status> : [message] <Message> が基本系
// Paramの情報を保持している場合はParamの要素毎に [<Paramのkey>] : <Paramのvalue> を追記して出力する
func (e *Error) Error() string {
	errMessage := fmt.Sprintf("[code] %v : [message] %v", e.errCode, e.message)

	for key, value := range e.param {
		errMessage = errMessage + fmt.Sprintf(" : [%v] %v", key, value)
	}

	return errMessage
}

// HasErrCode 保持しているエラーコードと引数に取ったエラーコードが同じかを判定
func (e *Error) HasErrCode(errCode uint16) bool {
	return e.errCode == errCode
}

// existsErrCode 指定したエラーコードが存在するかを判定
func existsErrCode(errCode uint16) bool {
	_, existsErrCode := errList[errCode]
	return existsErrCode
}

// newErr エラーコードとエラーメッセージを受け取ってErrorを生成
func newErr(errCode uint16, message string) *Error {
	return &Error{errCode: errCode, message: message}
}

// newParamErr エラーコードとエラーメッセージとパラメータ情報を受け取ってErrorを生成
func newParamErr(errCode uint16, message string, param map[string]string) *Error {
	return &Error{errCode: errCode, message: message, param: param}
}
