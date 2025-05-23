package services

import (
	"backend/clients"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("mi_clave_secreta_super_segura")

func Login(username string, password string) (string, error) {
	user := clients.GetUserByUsername(username)
	if user.Username == "" {
		return "", errors.New("usuario no encontrado")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("contraseña incorrecta")
	}

	// Generar el JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role, // 👈 agregamos el rol
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
