package requests

type AddressesRequest struct {
	UserID        int64   `json:"user_id" binding:"required"`
	RecipientName string  `json:"recipient_name" binding:"required"`
	PhoneNumber   string  `json:"phone_number" binding:"required"`
	StreetAddress string  `json:"street_address" binding:"required"`
	City          *string `json:"city" binding:"required"`
	Province      *string `json:"province" binding:"required"`
	PostalCode    *string `json:"postal_code" binding:"required"`
}
