package repository

import (
	"be_test_linkque/core/entity"
	"be_test_linkque/infrastructure/v1/mysql/models"

	"gorm.io/gorm"
)

type (
	UserRepository struct {
		DB *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) Create(req *entity.User) error {
	user := repo.mapEntityToModel(req)

	res := repo.DB.Create(&user)
	return res.Error
}

func (repo *UserRepository) mapEntityToModel(e *entity.User) *models.User {
	m := new(models.User)
	m.Name = e.Name
	m.Age = e.Age
	m.City = e.City
	return m
}
