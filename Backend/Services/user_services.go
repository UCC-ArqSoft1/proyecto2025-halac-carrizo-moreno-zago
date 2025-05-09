package services

import (
	"backend/clients"
	"fmt"
)

func Login(username string, password string) {
	//Get the user from the database
	user := clients.GetUserByUsername(username)
	fmt.Println("Usuario Obtenido: ", user)

	//Hash the password

	//Compare the hashed password with the stored password

	//If they match, generate a JWT token
}
