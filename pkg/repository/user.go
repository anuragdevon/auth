package repository

import (
	"auth/pkg/repository/models"
)

func (db *Database) CreateUser(user *models.User) (int, error) {
	result := db.DB.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.Id), nil
}

func (db *Database) GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (db *Database) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
