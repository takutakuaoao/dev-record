package category

import "github.com/fatih/color"

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c *Category) ToArray() []string {
	color.Blue("Prints %s in blue.", "text")
	return []string{c.Id, c.Name, c.Slug}
}
