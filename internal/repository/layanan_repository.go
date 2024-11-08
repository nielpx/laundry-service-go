package repository

import (
	"golang-gorm-gin/internal/models"
	"gorm.io/gorm"
)

type LayananRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id int) (*models.Product, error)
	Create(product *models.Product) error
	Update(id int, product *models.Product) error
	Delete(id int) error
}

type layananRepository struct{
	db *gorm.DB
}

func NewLayananRepository(db *gorm.DB) LayananRepository {
	return &layananRepository{db}
}

func (r *layananRepository) FindAll() ([]models.Product, error){
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}
func (r *layananRepository) FindByID(id int) (*models.Product, error){
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *layananRepository) Create(product *models.Product) error{
	return r.db.Create(product).Error
}

func (r *layananRepository) Update(id int, product *models.Product) error{
	return r.db.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *layananRepository) Delete(id int) error{
	return r.db.Delete(&models.Product{}, id).Error
}