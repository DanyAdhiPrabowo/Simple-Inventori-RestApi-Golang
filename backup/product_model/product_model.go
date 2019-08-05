package models

import (
	"database/sql"

	"../../entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			var category int
			err2 := rows.Scan(&id, &name, &price, &quantity, &category)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
					Category: category,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) FindSpecific(id int64) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where id=?", id)
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			var category int
			err2 := rows.Scan(&id, &name, &price, &quantity, &category)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
					Category: category,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Search(keyword string) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			var category int
			err2 := rows.Scan(&id, &name, &price, &quantity, &category)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
					Category: category,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Create(product *entities.Product) (err error) {
	result, err := productModel.Db.Exec("insert into product(name, price, quantity, id_category) values(?, ?, ?, ?)", product.Name, product.Price, product.Quantity, product.Category)
	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}
}

func (productModel ProductModel) Update(id int64, product *entities.ProductEdit) (int64, error) {
	result, err := productModel.Db.Exec("Update product  set name=?, price=?, quantity=?, id_category=? where id=?", product.Name, product.Price, product.Quantity, product.Category, id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("Delete from product where id=?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
