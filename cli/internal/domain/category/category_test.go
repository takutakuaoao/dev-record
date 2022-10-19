package category

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestCategory_ToArray(t *testing.T) {
	type fields struct {
		Id   string
		Name string
		Slug string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"正常系", fields{"test", "PHP", "php"}, []string{"test", "PHP", "php"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Category{
				Id:   tt.fields.Id,
				Name: tt.fields.Name,
				Slug: tt.fields.Slug,
			}
			testza.AssertEqual(t, c.ToArray(), tt.want)
		})
	}
}
