package main

import (
	"backend/clients"
	"backend/dao"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	createUser("admin", "admin123", "admin")
	createUser("lucia", "socio123", "socio")
}

func createUser(username, password, role string) {
	var existing dao.User
	result := clients.DB.First(&existing, "username = ?", username)
	if result.Error == nil {
		fmt.Printf("⚠️ Usuario %s ya existe. Saltando...\n", username)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := dao.User{
		Username:     username,
		PasswordHash: string(hash),
		Role:         role,
	}

	if err := clients.DB.Create(&user).Error; err != nil {
		fmt.Printf("⚠️ Error creando %s: %v\n", username, err)
	} else {
		fmt.Printf("✅ Usuario creado: %s (%s)\n", username, role)
	}
}

