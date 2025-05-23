package dao

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"not null"` // "admin" o "socio"
}
