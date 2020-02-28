package user

import (
	"math/big"
	"database/sql"
)

type User struct {
	TableName string

	Id struct {
		Name string
		Value big.Int
	}

	Name struct {
		Name string
		Value string
	}
}

func NewUser() User {
	user := User{}
	user.TableName = "users"
	user.Id.Name = "id"
	user.Name.Name = "name"
	return user
}

func (this *User) Migrate(db *sql.DB) {
	
}