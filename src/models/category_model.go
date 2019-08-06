package models

import (
	"database/sql"

	"../entities"
)

// CategoryModel adalah
type CategoryModel struct {
	Db *sql.DB
}

// FindAll = model untuk menampilkan semua data
func (categoryModel CategoryModel) FindAll() (category []entities.Category, err error) {
	rows, err := categoryModel.Db.Query("select * from category")
	// Jika Terjadi error query
	if err != nil {
		return nil, err
	}
	// Jika TIDAK Terjadi error query
	var categories []entities.Category
	for rows.Next() {
		var idCategory int64
		var nameCategory string
		err2 := rows.Scan(&idCategory, &nameCategory)
		// Jika Terjadi error saat scan value
		if err2 != nil {
			return nil, err2
		}
		// Jika TIDAK Terjadi error saat scan value
		category := entities.Category{
			ID:   idCategory,
			Name: nameCategory,
		}
		categories = append(categories, category)

	}
	return categories, nil

}

// FindSpecific = model untuk menampilkan data secara spesifik berdasarkan id category
func (categoryModel CategoryModel) FindSpecific(id int64) (category []entities.Category, err error) {
	rows, err := categoryModel.Db.Query("select * from category where id_category=?", id)
	// Jika Terjadi error query
	if err != nil {
		return nil, err
	}
	// Jika TIDAK Terjadi error query
	var categories []entities.Category
	for rows.Next() {
		var id int64
		var nameCategory string
		err2 := rows.Scan(&id, &nameCategory)
		// Jika Terjadi error saat scan value
		if err2 != nil {
			return nil, err2
		}
		// Jika TIDAK Terjadi error saat scan value
		category := entities.Category{
			ID:   id,
			Name: nameCategory,
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// Search = model untuk menampilkan semua data berdasarkan nama category yang diinput
func (categoryModel CategoryModel) Search(keyword string) (category []entities.Category, err error) {
	rows, err := categoryModel.Db.Query("select * from category where name_category like ?", "%"+keyword+"%")
	// Jika Terjadi error query
	if err != nil {
		return nil, err
	}
	// Jika TIDAK Terjadi error query
	var categories []entities.Category
	for rows.Next() {
		var idCategory int64
		var nameCategory string
		err2 := rows.Scan(&idCategory, &nameCategory)
		// Jika Terjadi error saat scan value
		if err2 != nil {
			return nil, err2
		}
		// Jika Terjadi error saat scan value
		category := entities.Category{
			ID:   idCategory,
			Name: nameCategory,
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// Create = model untuk menyimpan data baru
func (categoryModel CategoryModel) Create(category *entities.Category) (err error) {
	result, err := categoryModel.Db.Exec("insert into category (name_category) values (?)", category.Name)
	// Jika Terjadi error query
	if err != nil {
		return err
	}
	// Jika TIDAk Terjadi error query
	category.ID, _ = result.LastInsertId()
	return nil
}

// Update = model untuk mengupdate data berdasarkan id category
func (categoryModel CategoryModel) Update(id int64, category *entities.CategoryEdit) (int64, error) {
	result, err := categoryModel.Db.Exec("Update category  set name_category=? where id_category=?", category.Name, id)
	// Jika Terjadi error query
	if err != nil {
		return 0, err
	}
	// Jika TIDAK Terjadi error query
	return result.RowsAffected()
}

// Delete = model untuk menghapus data berdasarkan id category
func (categoryModel CategoryModel) Delete(id int64) (int64, error) {
	result, err := categoryModel.Db.Exec("Delete from category where id_category=?", id)
	// Jika Terjadi error query
	if err != nil {
		return 0, err
	}
	// Jika TIDAK Terjadi error query
	return result.RowsAffected()
}
