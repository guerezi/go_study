package models

type House struct {
	ID      int     `json:"id" bson:"id" validate:"required"`
	Street  string  `json:"street" bson:"street" validate:"required"`
	Number  int     `json:"number" bson:"number" validate:"required"`
	City    string  `json:"city" bson:"city" validate:"required"`
	State   string  `json:"state" bson:"state" validate:"required"`
	ZipCode string  `json:"zip_code" bson:"zip_code" validate:"required"`
	Price   float64 `json:"price" bson:"price" validate:"required,gte=0"`
	OwnerID *int    `json:"owner_id" bson:"owner_id"`
}

// se me atacar eu vou atacar
// TODO: criar um model de houses

// TODO: authenticaçao com sessão ??
// como faz pra verificar uma sessão mds
// if tem no banco user e password, deu boa :)
// Criar um novo serviço no docker (redis/varkey)

// Ter no redis um codigo bonito
// id do user, hash
// valor vai ser o token usado de verdade?

// sahdvkahjsdbjkajks : 123 ??
// 123 : { gabrielzinho : 123 }
// Sei lá porra

// redis pra cache
