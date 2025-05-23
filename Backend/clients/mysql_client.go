package clients

import (
	"backend/dao"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	user := "gym_admin"
	password := "admin123"
	host := "localhost"
	port := 3306
	database := "gym_db"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, password, host, port, database)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error connecting to DB: %v", err))
	}

	DB.AutoMigrate(&dao.User{})
}

func GetUserByUsername(username string) dao.User {
	var user dao.User
	DB.First(&user, "username = ?", username)
	return user
}
