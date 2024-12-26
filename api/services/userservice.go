package services

import "shopping-site/api/repositories"

type UserService interface {
}

type userService struct {
	repositories.UserRepository
}

func CommenceUserService(user repositories.UserRepository) UserService {
	return &userService{user}
}
