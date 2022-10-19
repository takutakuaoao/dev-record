package pkg

import (
	"dev-record-cli/internal/exception"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

// errors.WithStackのテスト
func TestPkgErrors_WithStack(t *testing.T) {
	err := errors.New("new Error")
	stackErr := errors.WithStack(err)
	stackTrace := errors.Errorf("error %+v", stackErr)

	// テストしやすいようにスタックトレースを改行毎に分割
	stackTraceList := strings.Split(stackTrace.Error(), "\n")

	// スタックトレースの一行目に error new Error となっていることのアサート
	if stackTraceList[0] != "error new Error" {
		t.Error("pkg/errors WithStack is having bug")
	}
}

func Test_WithStack_Multiple(t *testing.T) {
	t.Logf("%+v", causeError1())
}

func causeError1() error {
	return causeError2()
}

func causeError2() error {
	err := causeError3()

	return errors.Wrap(exception.NewFromErrCode(100), err.Error())
}

func causeError3() error {
	// return errors.WithStack(exception.NewParamErrFromCode(exception.ERR_API_URL_NOT_FOUND, map[string]string{"url": "xxx.xxx"}))
	return errors.Wrap(exception.NewParamErrFromCode(exception.ERR_API_URL_NOT_FOUND, map[string]string{"url": "xxx.xxx"}), "error")
}
