package entities

import "fmt"

type Category struct {
	Id   int64  `json:"id"` //json:"id" id adalah nama custom yang ditampilkan di result
	Name string `json:"name"`
}

type CategoryEdit struct {
	Name string `json:"name"`
}

func (category Category) toString() string {
	return fmt.Sprintf("id_category: %d\nname_category: %d", category.Id, category.Name)
}

func (categoryEdit CategoryEdit) toString() string {
	return fmt.Sprintf("name: %d", categoryEdit.Name)
}
