package repositories

import (
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/internal/app/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Db: db,
	}
}

func (u *ProductRepository) CreateProduct(prodcut *models.Product) error {
	tx := u.Db.Begin()

	// if the program panics due to any error, recover at the end and rollback the transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()
	// if transaction cannot begin then raise an error
	if err := tx.Error; err != nil {
		return fmt.Errorf("beginning transaction failed: %v", err)
	}

	if err := tx.Create(&prodcut).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("creating product failed: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commiting product failed: %v", err)
	}

	return nil
}

func (u *ProductRepository) DeleteProduct(productId int) error {
	tx := u.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if err := tx.Error; err != nil {
		return fmt.Errorf("beggining trancation failed: %v", err)
	}
	var product *models.Product

	if err := tx.Model(&models.Product{}).Where("id = ?", productId).First(&product).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("finding product by id failed: %v", err)
	}

	if err := tx.Delete(&product).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("deleting product failed: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commiting transaction failed: %v", err)
	}

	return nil
}

func (u *ProductRepository) UpdateProduct(product *models.Product) error {
	tx := u.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return fmt.Errorf("beginning transaction failed: %v", err)
	}

	if err := tx.Model(&models.Product{}).Where("id = ?", product.ID).Updates(
		map[string]interface{}{
			"name":        product.Name,
			"description": product.Description,
			"price":       product.Price,
		}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("updating prodcut failed: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commiting transaction failed: %v", err)
	}

	return nil
}

func (u *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var prodcut models.Product

	err := u.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&prodcut, id).Error; err != nil {
			return fmt.Errorf("finding product by id failed: %v", err)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("transaction failed: %v", err)
	}

	return &prodcut, nil
}
