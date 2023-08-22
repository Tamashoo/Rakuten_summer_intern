package repository

import (
	"backend/model"
	"errors"

	"gorm.io/gorm"
)

type IUserReduceRepository interface {
	GetUserReduceByUsername(userReduce *model.UserReduce, username string) error
	CreateUserReduce(userReduce *model.UserReduce) error
	UpdateUserReduce(userReduce *model.UserReduce) error
}

type userReduceRepository struct {
	db *gorm.DB
}

func NewUserReduceRepository(db *gorm.DB) IUserReduceRepository {
	return &userReduceRepository{db}
}

func (urr *userReduceRepository) GetUserReduceByUsername(userReduce *model.UserReduce, username string) error {
	if err := urr.db.Where("username = ?", username).First(userReduce).Error; err != nil {
		return err
	}
	return nil
}

func (urr *userReduceRepository) CreateUserReduce(userReduce *model.UserReduce) error {
	if err := urr.db.Create(userReduce).Error; err != nil {
		return err
	}
	return nil
}

func (urr *userReduceRepository) UpdateUserReduce(userReduce *model.UserReduce) error {
	result := urr.db.Model(&userReduce).Where("username = ?", userReduce.Username).
		Updates(map[string]interface{}{"exp": userReduce.Exp, "reduce_score": userReduce.ReduceScore})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
