package category

type Category struct {
	ID   		string  `json:"id" validate:"omitempty,uuid"`
	Name 		string 	`json:"name" validate:"required,min=3,max=100"`
	Description string 	`json:"description" validate:"required,min=10,max=500"`
}