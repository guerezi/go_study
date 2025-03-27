package models

type User struct {
	ID    int    `json:"id" bson:"id" validate:"required"`
	Name  string `json:"name" bson:"name" validate:"required"`
	Email string `json:"email" bson:"email" validate:"required,email"`
	Age   *int   `json:"age" bson:"age" validate:"required,gte=0,lte=130"`
}

type Houses struct {
	ID int `json:"id" bson:"id" validate:"required"`
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
