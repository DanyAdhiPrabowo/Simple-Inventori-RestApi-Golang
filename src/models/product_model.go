package models

import (
	"database/sql"

	"../entities"
)

// ProductModel adalah
type ProductModel struct {
	Db *sql.DB
}

// GetProduct = model untuk menampilkan semua data atau data berdasarkan id category
func (productModel ProductModel) GetProduct(idCategory int64) (product []entities.Product, err error) {
	// Jika value id_category ada
	if idCategory != 0 {
		rows, err := productModel.Db.Query("SELECT a.id, a.name, a.price, a.quantity, b.name_category FROM `product` AS `a` LEFT JOIN `category` AS `b` ON `a`.`id_category`=`b`.`id_category` where a.id_category=?", idCategory)
		// Jika Terjadi error query
		if err != nil {
			return nil, err
		}
		// Jika TIDAK Terjadi error query
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int
			var category string
			err2 := rows.Scan(&id, &name, &price, &quantity, &category)
			// Jika Terjadi error saat scan value
			if err2 != nil {
				return nil, err2
			}
			// Jika TIDAK Terjadi error saat scan value
			product := entities.Product{
				ID:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
				Category: category,
			}
			products = append(products, product)
		}
		return products, nil

	}
	// Jika value id_category TIDAK ada
	rows, err := productModel.Db.Query("SELECT a.id, a.name, a.price, a.quantity, b.name_category FROM `product` AS `a` LEFT JOIN `category` AS `b` ON `a`.`id_category`=`b`.`id_category`")
	// Jika Terjadi error query
	if err != nil {
		return nil, err
	}
	// Jika TIDAK Terjadi error query
	var products []entities.Product
	for rows.Next() {
		var id int64
		var name string
		var price float64
		var quantity int
		var category string
		err2 := rows.Scan(&id, &name, &price, &quantity, &category)
		// Jika Terjadi error saat scan value
		if err2 != nil {
			return nil, err2
		}
		// Jika TIDAk Terjadi error saat scan value
		product := entities.Product{
			ID:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Category: category,
		}
		products = append(products, product)
	}
	return products, nil
}

// FindSpecific = model untuk menampilkan data secara spesifik berdasarkan id category
func (productModel ProductModel) FindSpecific(id int64) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("SELECT a.id, a.name, a.price, a.quantity, b.name_category FROM `product` AS `a` LEFT JOIN `category` AS `b` ON `a`.`id_category`=`b`.`id_category` where a.id=?", id)
	// Jika Terjadi error query
	if err != nil {
		return nil, err
	}
	// Jika TIDAK Terjadi error query
	var products []entities.Product
	for rows.Next() {
		var id int64
		var name string
		var price float64
		var quantity int
		var category string
		err2 := rows.Scan(&id, &name, &price, &quantity, &category)
		// Jika Terjadi error saat scan value
		if err2 != nil {
			return nil, err2
		}
		// Jika TIDAK Terjadi error saat scan value
		product := entities.Product{
			ID:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Category: category,
		}
		products = append(products, product)
	}
	return products, nil
}

// Search = model untuk menampilkan semua data berdasarkan nama category yang diinput
func (productModel ProductModel) Search(keyword string) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("SELECT a.id, a.name, a.price, a.quantity, b.name_category FROM `product` AS `a` LEFT JOIN `category` AS `b` ON `a`.`id_category`=`b`.`id_category` where a.name like ?", "%"+keyword+"%")
	// Jika Terjadi error query
	if err != nil {
		return nil, err
	}
	// Jika Terjadi error query
	var products []entities.Product
	for rows.Next() {
		var id int64
		var name string
		var price float64
		var quantity int
		var category string
		err2 := rows.Scan(&id, &name, &price, &quantity, &category)
		// Jika Terjadi error saat scan value
		if err2 != nil {
			return nil, err2
		}
		// Jika TIDAK Terjadi error saat scan value
		product := entities.Product{
			ID:       id,
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Category: category,
		}
		products = append(products, product)
	}
	return products, nil
}

// Create = model untuk menyimpan data baru
func (productModel ProductModel) Create(product *entities.Product) (err error) {
	result, err := productModel.Db.Exec("insert into product(name, price, quantity, id_category) values(?, ?, ?, ?)", product.Name, product.Price, product.Quantity, product.Category)
	// Jika Terjadi error query
	if err != nil {
		return err
	}
	// Jika TIDAK Terjadi error query
	product.ID, _ = result.LastInsertId()
	return nil
}

// Update = model untuk mengupdate data berdasarkan id category
func (productModel ProductModel) Update(id int64, product *entities.ProductEdit) (int64, error) {
	result, err := productModel.Db.Exec("Update product  set name=?, price=?, quantity=?, id_category=? where id=?", product.Name, product.Price, product.Quantity, product.Category, id)
	// Jika Terjadi error query
	if err != nil {
		return 0, err
	}
	// Jika TIDAK Terjadi error query
	return result.RowsAffected()
}

// Delete = model untuk menghapus data berdasarkan id category
func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("Delete from product where id=?", id)
	// Jika Terjadi error query
	if err != nil {
		return 0, err
	}
	// Jika TIDAk Terjadi error query
	return result.RowsAffected()
}
