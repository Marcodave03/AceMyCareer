package products

import (
	"database/sql"

	_ "github.com/lib/pq"
	"log"
)

func CreateTableProducts(db *sql.DB) {
	query := `
        CREATE SCHEMA IF NOT EXISTS Products;
        CREATE TABLE IF NOT EXISTS Products.Products (
            product_id SERIAL PRIMARY KEY,
            product_name VARCHAR(30),
            product_image_link VARCHAR(30),
            product_desc TEXT,
            product_price INT,
            product_is_available BOOL,
            product_type_id INT REFERENCES Products.ProductTypes(product_type_id) ON DELETE SET NULL
        );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func DeleteProduct(db *sql.DB, request DeleteProductRequest) error {
	query := `
    DELETE FROM Products.Products WHERE Products.Products.product_id = $1;
    `
	_, err := db.Exec(query, request.ProductId)
	return err
}

func CreateTableProductTypes(db *sql.DB) {
	query := `
        CREATE SCHEMA IF NOT EXISTS Products;
        CREATE TABLE IF NOT EXISTS Products.ProductTypes (
            product_type_id SERIAL PRIMARY KEY,
            product_type_name VARCHAR(30),
            product_type_desc TEXT,
            product_type_discount INT,
            product_category_id INT REFERENCES Products.ProductCategories(product_category_id)
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func CreateTableProductCategories(db *sql.DB) {
	query := `
        CREATE SCHEMA IF NOT EXISTS Products;
        CREATE TABLE IF NOT EXISTS Products.ProductCategories (
            product_category_id INT PRIMARY KEY,
            product_category_name VARCHAR(30),
            product_category_desc TEXT
        );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func CreateProduct(db *sql.DB, newProduct Product) error {
	query := `
    INSERT INTO Products.Products
    (
        product_name,
        product_image_link,
        product_desc,
        product_price,
        product_is_available,
        product_type_id
    )
    VALUES ( $1, $2, $3, $4, $5, $6);
    `
	_, err := db.Exec(
		query,
		newProduct.ProductName,
		newProduct.ProductImageLink,
		newProduct.ProductDesc,
		newProduct.ProductPrice,
		newProduct.ProductIsAvailable,
		newProduct.ProductTypeId,
	)
	return err
}

func GetAllProducts(db *sql.DB) ([]Product, error) {
	query := `SELECT * FROM Products.Products;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.ProductImageLink,
			&product.ProductDesc,
			&product.ProductPrice,
			&product.ProductIsAvailable,
			&product.ProductTypeId,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, rows.Err()
}

func DropTableProducts(db *sql.DB) {
	query := `DROP TABLE IF EXISTS Products.Products;`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetAllProductTypes(db *sql.DB) ([]ProductType, error) {
	query := `SELECT * FROM Products.ProductTypes;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productTypes []ProductType
	for rows.Next() {
		var newPorductType ProductType
		err := rows.Scan(
			&newPorductType.ProductTypeId,
			&newPorductType.ProductTypeName,
			&newPorductType.ProductTypeDesc,
			&newPorductType.ProductTypeDiscount,
			&newPorductType.ProductCategoryId,
		)
		if err != nil {
			return nil, err
		}
		productTypes = append(productTypes, newPorductType)
	}
    if err := rows.Err(); err != nil {
		return nil, err
	}
	return productTypes, nil
}

func DropTableProductTypes(db *sql.DB) {
    _, err := db.Exec( "DROP TABLE IF EXISTS Products.ProductTypes;",)
    if err != nil {
        log.Fatal(err.Error())
    }
}

func CreateProductType(db *sql.DB, productTypeInfo ProductType) error {
	query := `
    INSERT INTO Porducts.ProductTypes
    (
    product_type_id,
    product_type_name,
    product_type_desc,
    product_type_discount,
    product_category_id,
    )
    VALUES ( $1, $2, $3, $4, $5);
    `
	_, err := db.Exec(query,
		productTypeInfo.ProductTypeId,
		productTypeInfo.ProductTypeName,
		productTypeInfo.ProductTypeDesc,
		productTypeInfo.ProductTypeDiscount,
		productTypeInfo.ProductCategoryId,
	)
	return err
}
