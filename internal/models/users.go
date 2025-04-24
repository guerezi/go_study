package models

type User struct {
	ID           int    `json:"id" bson:"id" validate:"required"`
	Name         string `json:"name" bson:"name" validate:"required"`
	Email        string `json:"email" bson:"email" validate:"required,email"`
	Age          *int   `json:"age" bson:"age" validate:"gte=0,lte=130"`
	PasswordHash string `json:"password" bson:"password"`
}

// TODO
// type password struct {
// 	Password string `json:"password" bson:"password" validate:"required"`
// 	Hash     string `json:"hash" bson:"hash" validate:"required"`
// }

// func (p* password) HashPassword() (string, error) {
// 	// ir e voltar de hash
// }
