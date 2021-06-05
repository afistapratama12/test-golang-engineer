package user

import "gorm.io/gorm"

type Repository interface {
	FindByEmail(email string) (User, error)
	Create(data User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(data User) (User, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
