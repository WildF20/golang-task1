package model

type Category struct {
	ID   		string  `json:"id" validate:"omitempty"`
	Name 		string 	`json:"name" validate:"required, min=3,max=100"`
	Description string 	`json:"description" validate:"required, min=10,max=500"`
}