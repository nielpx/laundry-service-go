package usecase

import (
	"golang-gorm-gin/internal/models"
	"golang-gorm-gin/internal/repository"

	"github.com/sirupsen/logrus"
)

type LayananUsecase interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	CreateLayanan(product *models.Product) error
	UpdateLayanan(id int, product *models.Product) error
	DeleteLayanan(id int) error
}

type layananRepository struct{
	layananRepo repository.LayananRepository
}

func NewLayananUsecase(repo repository.LayananRepository) LayananUsecase{
	return &layananRepository{repo}
}

func (u *layananRepository) GetAll() ([]models.Product, error){
	logrus.Info("Executing usecase GetAll function")
	return u.layananRepo.FindAll()
}

func (u *layananRepository) GetByID(id int) (*models.Product, error){
	logrus.Info("Executing usecase GetByID function")
	return u.layananRepo.FindByID(id)
}
func (u *layananRepository) CreateLayanan(product *models.Product) error{
	logrus.Info("Executing usecase CreateLayanan")
	return u.layananRepo.Create(product)
}

func (u *layananRepository) UpdateLayanan(id int, product *models.Product) error{
	logrus.Info("Executing usecase UpdateLayanan")
	return u.layananRepo.Update(id, product)
}

func (u *layananRepository) DeleteLayanan(id int) error {
	logrus.Info("Executing usecase DeleteLayanan")
    return u.layananRepo.Delete(id)
}