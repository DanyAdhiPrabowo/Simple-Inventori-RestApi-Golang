package entities

import "fmt"

// Category berfungsi untuk mengatur response setelah input data atau saat menampilkan data
type Category struct {
	ID   int64  `json:"id"` //json:"id" id adalah nama custom yang ditampilkan di result
	Name string `json:"name"`
}

// CategoryEdit berfungsi untuk mengatur response setelah mengedit data
type CategoryEdit struct {
	Name string `json:"name"`
}

func (category Category) toString() string {
	return fmt.Sprintf("id_category: %d\nname_category: %s", category.ID, category.Name)
}

func (categoryEdit CategoryEdit) toString() string {
	return fmt.Sprintf("name: %s", categoryEdit.Name)
}
