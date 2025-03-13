package models

type User struct {
	ID    int    `json:"id" bson:"id" validate:"required"`
	Name  string `json:"name" bson:"name" validate:"required"`
	Email string `json:"email" bson:"email" validate:"required,email"`
	Age   *int   `json:"age" bson:"age" validate:"required,gte=0,lte=130"`
}
