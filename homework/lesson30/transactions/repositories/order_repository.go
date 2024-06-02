package repositories

import (
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/internal/app/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (o *OrderRepository) Create(userId, productId int) error {
	var (
		user    *models.User
		prodcut *models.Product
	)

	tx := o.Db.Begin()
	if err := tx.Error; err != nil {
		return fmt.Errorf("beginning transaction for order failed: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if err := tx.Model(&user).Where("id = ?", userId).First(&user).Error; err != nil {
		return fmt.Errorf("finding user failed: %v", err)
	}

	if err := tx.Model(&prodcut).Where("id = ?", productId).First(&prodcut).Error; err != nil {
		return fmt.Errorf("finding product failed: %v", err)
	}

	if err := tx.Create(&models.Order{
		UserId:    userId,
		ProductId: productId,
		User:      *user,
		Product:   *prodcut,
	}).Error; err != nil {
		return fmt.Errorf("creating order failed: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commiting transaction failed: %v", err)
	}
	return nil
}

func (o *OrderRepository) GetOrderById(orderId int) (*models.Order, error) {
	var order models.Order

	tx := o.Db.Begin()
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("beginning order transaction failed: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Preload("User").Preload("Product").Where("id = ?", orderId).First(&order).Error; err != nil {
		return nil, fmt.Errorf("finding order failed: %v", err)
	}

	// Only commit the transaction if everything was successful
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("committing order transaction failed: %v", err)
	}

	return &order, nil
}
