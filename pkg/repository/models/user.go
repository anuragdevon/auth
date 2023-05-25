package models

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}
