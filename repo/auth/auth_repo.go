package auth

import (
	"rest/entities"

	"gorm.io/gorm"
)

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

type AuthRepo struct {
	db *gorm.DB
}

func (authRepo AuthRepo) Login(user entities.User) (entities.User, error) {
	userDb := FromEntities(user)
	result := authRepo.db.First(&userDb, "email = ? AND password = ?", userDb.Email, userDb.Password)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return userDb.ToEntities(), nil
}
