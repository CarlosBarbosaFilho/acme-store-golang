package model

import "curso-alura-golang-I/db"

/*
*
Criando estrutura de produtos
*
*/
type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amout       int
}

func ListAllProducts() []Product {

	db := db.ConectionDB()
	selectAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}
	for selectAllProducts.Next() {
		var id int
		var name, description string
		var price float64
		var amount int

		err = selectAllProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amout = amount

		products = append(products, product)
	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, amount int) {
	db := db.ConectionDB()
	insertProduct, err := db.Prepare("insert into products(name, description, price, amount) values ($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, amount)

	defer db.Close()
}

func RemoveProduct(id string) {
	db := db.ConectionDB()
	deleteProduct, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func Edit(id string) Product {

	db := db.ConectionDB()
	updateProduct, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productEditable := Product{}

	for updateProduct.Next() {
		var id int
		var name, description string
		var price float64
		var amount int

		err = updateProduct.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		productEditable.Id = id
		productEditable.Name = name
		productEditable.Description = description
		productEditable.Price = price
		productEditable.Amout = amount

	}

	defer db.Close()
	return productEditable
}

func Update(id int, name, description string, price float64, amount int) {

	db := db.ConectionDB()
	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, amount, id)
	db.Close()
}
