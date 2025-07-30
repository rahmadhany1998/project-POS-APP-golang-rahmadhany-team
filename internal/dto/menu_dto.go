package dto

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	ItemCode    string `json:"item_code" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	RetailPrice int    `json:"retail_price"`
	Photo       string `json:"photo"`
	Stock       int    `json:"stock"`
	CategoryID  uint   `json:"category_id" validate:"required"`
	Available   bool   `json:"available"`
	Quantity    int    `json:"quantity"`
	Unit        string `json:"unit"`
	Status      string `json:"status"`
}

type MenuFilter struct {
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
	Available  *bool  `json:"available"`
	Status     string `json:"status"`
	MinPrice   *int   `json:"min_price"`
	MaxPrice   *int   `json:"max_price"`
}
