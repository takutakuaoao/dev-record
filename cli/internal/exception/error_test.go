package exception

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestError_NewParamErrorFromCode(t *testing.T) {
	type field struct {
		ErrorCode uint16
		Param     map[string]string
	}

	tests := []struct {
		testCase string
		field    field
		want     string
	}{
		{
			testCase: "エラーコードとパラメータからErrorを生成するテスト",
			field:    field{ErrorCode: ERR_API_CONNECTION, Param: map[string]string{"name": "xxx"}},
			want:     "[code] 100 : [message] API Connection Error : [name] xxx",
		},
		{
			testCase: "エラーコードとパラメータからError生成時に存在しないエラーコードを指定した場合のテスト",
			field:    field{ErrorCode: 0, Param: map[string]string{"name": "xxx"}},
			want:     "[code] 1 : [message] Failed Make Error Struct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			err := NewParamErrFromCode(tt.field.ErrorCode, tt.field.Param)
			testza.AssertEqual(t, tt.want, err.Error())
		})
	}
}

func TestError_Error(t *testing.T) {
	tests := []struct {
		testCase    string
		errorStruct Error
		want        string
	}{
		{"パラメータなし", *NewFromErrCode(ERR_API_CONNECTION), "[code] 100 : [message] API Connection Error"},
		{"パラメータが1つ", *NewParamErrFromCode(ERR_API_CONNECTION, map[string]string{"url": "http://xxx.xxx/yyy"}), "[code] 100 : [message] API Connection Error : [url] http://xxx.xxx/yyy"},
		{"パラメータが2つ", *NewParamErrFromCode(ERR_API_CONNECTION, map[string]string{"ID": "xxx", "name": "yyy"}), "[code] 100 : [message] API Connection Error : [ID] xxx : [name] yyy"},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			testza.AssertEqual(t, tt.want, tt.errorStruct.Error())
		})
	}
}

func TestError_NewFromCode(t *testing.T) {
	tests := []struct {
		testCase  string
		errorCode uint16
		want      string
	}{
		{"ErrorCode: 100", 100, "[code] 100 : [message] API Connection Error"},
		{"ErrorCode: 101", 101, "[code] 101 : [message] API URL Not Found Error"},
		{"存在しないエラーコードを指定した場合にFAILED_MAKE_ERROR のエラーを返す", 0, "[code] 1 : [message] Failed Make Error Struct"},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			err := NewFromErrCode(tt.errorCode)
			testza.AssertEqual(t, tt.want, err.Error())
		})
	}
}

func TestError_IsErrorCode(t *testing.T) {
	type fields struct {
		HasCode     uint16
		ConfirmCode uint16
	}
	tests := []struct {
		testCase string
		fields   fields
		want     bool
	}{
		{
			testCase: "ErrCodeが100でErrCodeが100かを判定しtrueとなる",
			fields:   fields{HasCode: ERR_API_CONNECTION, ConfirmCode: ERR_API_CONNECTION},
			want:     true,
		},
		{
			testCase: "ErrCodeが100でErrCodeが101かを判定しfalseとなる",
			fields:   fields{HasCode: ERR_API_CONNECTION, ConfirmCode: ERR_API_URL_NOT_FOUND},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			err := NewFromErrCode(tt.fields.HasCode)
			testza.AssertEqual(t, err.HasErrCode(tt.fields.ConfirmCode), tt.want)
		})
	}
}
