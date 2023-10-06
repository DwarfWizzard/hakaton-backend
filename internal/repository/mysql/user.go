package mysql

import (
	"github.com/DwarfWizzard/hakaton-backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepo struct {
	dbClient *gorm.DB
}

func NewUserRepo(dbClient *gorm.DB) *UserRepo {
	return &UserRepo{
		dbClient: dbClient,
	}
}

func (r *UserRepo) GetUserByLogin(login string) (*domain.User, error) {
	var user domain.User
	result := r.dbClient.
		Table("users").
		Where(`users.name = ? `, login).
		Preload("StudentClass").
		Preload("TeacherClasses").
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepo) GetUserById(userId uint64) (*domain.User, error) {
	var user domain.User
	result := r.dbClient.
		Table("users").
		Where(`users.id = ? `, userId).
		Preload("StudentClass").
		Preload("TeacherClasses").
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
