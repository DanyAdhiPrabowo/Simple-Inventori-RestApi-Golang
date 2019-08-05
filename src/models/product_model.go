package models

import (
	"database/sql"

	"../entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) GetProduct(id_category int64) (product []entities.Product, err error) {
	// log.Printf("idnya adalah %v", id_category)
	if id_category != 0 {
		rows, err := productModel.Db.Query("SELECT a.id, a.name, a.price, a.quantity, b.name_category FROM `product` AS `a` LEFT JOIN `category` AS `b` ON `a`.`id_category`=`b`.`id_category` where a.id_category=?", id_category)
		if err != nil {
			return nil, err
		} else {
			var products []entities.Product
			for rows.Next() {
				var id int64
				var name string
				var price float64
				var quantity int
				var category string
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
	} else {
		rows, err := productModel.Db.Query("SELECT a.id, a.name, a.price, a.quantity, b.name_category FROM `product` AS `a` LEFT JOIN `category` AS `b` ON `a`.`id_category`=`b`.`id_category`")
		if err != nil {
			return nil, err
		} else {
			var products []entities.Product
			for rows.Next() {
				var id int64
				var name string
				var price float64
				var quantity int
				var category string
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
			var category string
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
			var category string
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
