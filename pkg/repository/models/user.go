package models

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string `gorm:"column:password"`
}
