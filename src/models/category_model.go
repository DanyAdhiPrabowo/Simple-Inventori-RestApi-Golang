package models

import (
	"database/sql"

	"../entities"
)

type CategoryModel struct {
	Db *sql.DB
}

func (categoryModel CategoryModel) FindAll() (category []entities.Category, err error) {
	rows, err := categoryModel.Db.Query("select * from category")
	if err != nil {
		return nil, err
	} else {
		var categories []entities.Category
		for rows.Next() {
			var id_category int64
			var name_category string
			err2 := rows.Scan(&id_category, &name_category)
			if err2 != nil {
				return nil, err2
			} else {
				category := entities.Category{
					Id:   id_category,
					Name: name_category,
				}
				categories = append(categories, category)
			}
		}
		return categories, nil
	}
}

func (categoryModel CategoryModel) FindSpecific(id int64) (category []entities.Category, err error) {
	rows, err := categoryModel.Db.Query("select * from category where id_category=?", id)
	if err != nil {
		return nil, err
	} else {
		var categories []entities.Category
		for rows.Next() {
			var id int64
			var name_category string
			err2 := rows.Scan(&id, &name_category)
			if err2 != nil {
				return nil, err2
			} else {
				category := entities.Category{
					Id:   id,
					Name: name_category,
				}
				categories = append(categories, category)
			}
		}
		return categories, nil
	}
}

func (categoryModel CategoryModel) Search(keyword string) (category []entities.Category, err error) {
	rows, err := categoryModel.Db.Query("select * from category where name_category like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		var categories []entities.Category
		for rows.Next() {
			var id_category int64
			var name_category string
			err2 := rows.Scan(&id_category, &name_category)
			if err2 != nil {
				return nil, err2
			} else {
				category := entities.Category{
					Id:   id_category,
					Name: name_category,
				}
				categories = append(categories, category)
			}
		}
		return categories, nil
	}
}

func (categoryModel CategoryModel) Create(category *entities.Category) (err error) {
	result, err := categoryModel.Db.Exec("insert into category (name_category) values (?)", category.Name)
	if err != nil {
		return err
	} else {
		category.Id, _ = result.LastInsertId()
		return nil
	}
}

func (categoryModel CategoryModel) Update(id int64, category *entities.CategoryEdit) (int64, error) {
	result, err := categoryModel.Db.Exec("Update category  set name_category=? where id_category=?", category.Name, id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (categoryModel CategoryModel) Delete(id int64) (int64, error) {
	result, err := categoryModel.Db.Exec("Delete from category where id_category=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
