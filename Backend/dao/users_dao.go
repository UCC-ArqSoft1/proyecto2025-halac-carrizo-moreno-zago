package dao

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"unique"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"not null"` // "admin" o "socio"
}
