package model

type User struct {
	Id       int64  `gorm:"primarykey" json:"id"`
	UserName string `gorm:"varchar(255)" json:"user_name"`
	Password string `gorm:"varchar(255)" json:"password"`
}
