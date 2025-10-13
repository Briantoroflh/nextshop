package users

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
	"nextshop/entities"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(user *entities.Users) error {
	return r.DB.Create(user).Error
}

func (r *Repository) GetUserByID(id uint64) (*entities.Users, error) {
	var user entities.Users
	err := r.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *Repository) UpdateUser(user *entities.Users) error {
	return r.DB.Save(user).Error
}

func (r *Repository) DeleteUser(id uint64) error {
	log.Printf("Deleting user with ID: %d", id)
	return r.DB.Model(&entities.Users{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
	}).Error
}

func (r *Repository) GetAllUsers() (*entities.Users, error) {
	var user *entities.Users
	err := r.DB.Find(&user).Error
	return user, err
}

func (r *Repository) GetUserByRole(role *entities.Users) (*entities.Users, error) {
	var user entities.Users

	err := r.DB.Find(&user).Where("role = ?", role).Error
	return &user, err
}