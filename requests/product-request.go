package requests

type ProductRequest struct {
	SellerID    int64           `json:"seller_id" binding:"required"`
	CategoryID  int64          `json:"category_id" binding:"required"`
	Name        string          `json:"name" binding:"required"`
	Slug        string          `json:"slug" binding:"required"`
	Description *string         `json:"description" binding:"required"`
	Price       float64         `json:"price" binding:"required"`
	Stock       int             `json:"stock" binding"required"`
}