/*
CategoryAPIは外部APIからのカテゴリ情報のCRUDを制御する
*/
package categoryapi

import (
	"dev-record-cli/internal"
	"dev-record-cli/internal/domain/category"
	"dev-record-cli/internal/exception"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type CategoryApi struct {
	baseURL string
}

type JsonData struct {
	Result bool  `json:"result"`
	Data   Items `json:"data"`
}

type Items struct {
	Items []category.Category `json:"items"`
}

var (
	ErrApiConnection = errors.New("API Connection is failed.")
)

func NewCategoryApi(baseURL string) *CategoryApi {
	return &CategoryApi{
		baseURL: baseURL,
	}
}

func (c *CategoryApi) FetchAll() ([]category.Category, error) {
	// カテゴリ一覧を取得（API）
	requestURL := c.baseURL + internal.API_URL_CATEGORY_INDEX
	res, err := http.Get(requestURL)

	if res.StatusCode != 200 || err != nil {
		return nil, errors.WithStack(exception.NewParamErrFromCode(exception.ERR_API_CONNECTION, map[string]string{"url": requestURL}))
	}

	defer res.Body.Close()

	// JSONに変換
	body, _ := ioutil.ReadAll(res.Body)

	// JSON => struct に変換
	var jsonResponse JsonData
	json.Unmarshal(body, &jsonResponse)

	return jsonResponse.Data.Items, nil
}
