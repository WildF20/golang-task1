package product

type Product struct {
	ID   		string  `json:"id" validate:"omitempty,uuid"`
	Name 		string 	`json:"name" validate:"required,min=3,max=100"`
	CategoryID  string 	`json:"category_id" validate:"required,uuid"`
	Price 		int 	`json:"price" validate:"required"`
	Stock 		int 	`json:"stock" validate:"required"`
}