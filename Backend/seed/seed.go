package main

import (
	"backend/clients"
	"backend/dao"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	clients.DB.Create(&dao.User{
		Username:     "admin",
		PasswordHash: string(hash),
	})
}
