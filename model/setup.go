package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/cobaa"))

	if err != nil {
		panic("Koneksi Error")
	}

	db.AutoMigrate(&Adzan{})
	db.AutoMigrate(&Kota{})
	db.AutoMigrate(&User{})

	DB = db
}
