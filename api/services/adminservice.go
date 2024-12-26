package services

import (
	"shopping-site/api/repositories"
	"shopping-site/pkg/models"
	"shopping-site/utils/dto"
)

type AdminService interface {
	AddCategoreyService(*models.Categories) *dto.ErrorResponse
	AddBrandService(*models.Brands) *dto.ErrorResponse
}

type adminService struct {
	repositories.AdminRepository
}

func CommenceAdminService(admin repositories.AdminRepository) AdminService {
	return &adminService{admin}
}

func (repo *adminService) AddCategoreyService(category *models.Categories) *dto.ErrorResponse {
	return repo.AddCategoreyRepository(category)
}

func (repo *adminService) AddBrandService(brand *models.Brands) *dto.ErrorResponse {
	return repo.AddBrandRepository(brand)
}
