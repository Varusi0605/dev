package repositories

import "gorm.io/gorm"

type UserRepository interface {
}

type userRepository struct {
	*gorm.DB
}

func CommenceUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
