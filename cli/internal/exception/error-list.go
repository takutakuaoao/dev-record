package exception

// エラーコード
const (
	FAILED_MAKE_ERROR     = 1
	ERR_API_CONNECTION    = 100
	ERR_API_URL_NOT_FOUND = 101
)

// エラーコードに対応するメッセージリスト
var errList = map[uint16]string{
	FAILED_MAKE_ERROR:     "Failed Make Error Struct",
	ERR_API_CONNECTION:    "API Connection Error",
	ERR_API_URL_NOT_FOUND: "API URL Not Found Error",
}
