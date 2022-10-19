package categoryapi

import (
	"dev-record-cli/internal"
	"dev-record-cli/internal/domain/category"
	"dev-record-cli/internal/exception"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pkg/errors"
	"gopkg.in/h2non/gock.v1"
)

func TestCategoryApi_FetchAll(t *testing.T) {
	defer gock.Off()

	want := []category.Category{
		{Id: "test-1", Name: "PHP", Slug: "php"},
		{Id: "test-2", Name: "Laravel", Slug: "laravel"},
		{Id: "test-3", Name: "WEB開発", Slug: "web-develop"},
	}
	mockResponse := JsonData{
		Result: true,
		Data: Items{
			Items: want,
		},
	}

	baseURL, _, _ := makeMockServerForCategoryIndexAPI(200, mockResponse)
	got, _ := NewCategoryApi(baseURL).FetchAll()
	testza.AssertEqual(t, want, got)
}

func TestCategoryApi_FetchAll_WhenError(t *testing.T) {
	defer gock.Off()

	type failedResponse struct {
		Result bool
	}
	baseURL, _, _ := makeMockServerForCategoryIndexAPI(500, failedResponse{Result: false})

	_, err := NewCategoryApi(baseURL).FetchAll()

	causeErr := errors.Cause(err)
	switch e := causeErr.(type) {
	case *exception.Error:
		testza.AssertTrue(t, e.HasErrCode(exception.ERR_API_CONNECTION))
	default:
		testza.AssertTrue(t, false, "エラーの型が exception.Error ではないのでテスト失敗")
	}

}

// makeMockServerForCategoryIndexAPI
// /api/categories のモックサーバーを作成する
// ResponseTypeはモックサーバーにアクセスした際に返却されるレスポンスの型を指定する
func makeMockServerForCategoryIndexAPI[ResponseType any](status int, mockResponse ResponseType) (domain string, subURL string, response ResponseType) {
	baseURL := "http://app"
	requestURL := internal.API_URL_CATEGORY_INDEX

	makeMockServer(baseURL, requestURL, status, mockResponse)

	return baseURL, requestURL, mockResponse
}

// makeMockServer
// モックサーバーを作成する
func makeMockServer(baseURL string, subURL string, status int, response any) {
	gock.New(baseURL).
		Get(subURL).
		Reply(status).
		JSON(response)
}
