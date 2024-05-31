package products

type Product struct {
	ProductId          int    `json:"prodcut_id"`
	ProductName        string `json:"product_name"`
	ProductDesc        string `json:"product_desc"`
	ProductPrice       int    `json:"product_price"`
	ProductImageLink   string `json:"product_image_link"`
	ProductIsAvailable bool   `json:"product_is_available"`
	ProductTypeId      *int    `json:"product_type_id"`
}
type DeleteProductRequest struct {
	ProductId          int    `json:"prodcut_id"`
}

type ProductType struct {
	ProductTypeId       int    `json:"product_type_id"`
	ProductTypeName     string `json:"product_type_name"`
	ProductTypeDesc     string `json:"product_type_desc"`
	ProductTypeDiscount int    `json:"product_type_discount"`
	ProductCategoryId   int    `json:"product_category_id"`
}

type ProductCategory struct {
	ProductCategoryId   int    `json:"product_category_id"`
	ProductCategoryName string `json:"product_category_name"`
	ProductCategoryDesc string `json:"product_category_desc"`
}


