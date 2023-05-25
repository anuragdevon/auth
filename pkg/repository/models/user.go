package models

type User struct {
	Id       int64  `gorm:"column:email;primaryKey"`
	Email    string `gorm:"column:email;not null;unique"`
	Password string `gorm:"column:password;not null"`
	Usertype string `gorm:"column:usertype;not null"`
}
