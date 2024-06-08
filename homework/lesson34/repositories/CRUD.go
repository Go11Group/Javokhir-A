package repositories

import (
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson34/internal/app/models"
	"gorm.io/gorm"
)

type UniRepo struct {
	*UserRepository
	*ProductRepository
	*OrderRepository
	*UniverseRepository
}

type Filter struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Where  string `json:"where"`
}

type UniverseRepositor struct {
	Db *gorm.DB
}

func NewUniRepo(db *gorm.DB) *UniRepo {
	userRepo := NewUserRepository(db)
	productRepo := NewProductRepository(db)
	orderRepo := NewOrderRepository(db)
	uniRepo := NewUniverseRepository(db)
	u := UniRepo{
		userRepo,
		productRepo,
		orderRepo,
		uniRepo,
	}
	return &u
}

type UniverseRepository struct {
	Db *gorm.DB
}

func NewUniverseRepository(db *gorm.DB) *UniverseRepository {
	return &UniverseRepository{
		Db: db,
	}
}

func (f *UniverseRepository) FetchAll(result interface{}, filter Filter) error {
	var tableName, columnName string

	switch result.(type) {
	case *[]models.User:
		tableName = (&models.User{}).TableName()
		columnName = "user_name"
	case *[]models.Product:
		tableName = (&models.Product{}).TableName()
		columnName = "product_name"
	case *[]models.Order:
		tableName = (&models.Order{}).TableName()
		columnName = "order_name"
	default:
		return fmt.Errorf("invalid model type")
	}

	query := fmt.Sprintf("%s LIKE ?", columnName)
	fmt.Println(tableName)

	if err := f.Db.Where(query, "%%"+filter.Where+"%%").Limit(filter.Limit).Offset(filter.Offset).Find(result).Error; err != nil {
		return fmt.Errorf("failed to fetch all records: %v", err)
	}

	return nil
}

func (u *UniverseRepository) Create(model interface{}) error {
	tableName := ""

	switch model.(type) {
	case *models.User:
		tableName = (&models.User{}).TableName()
	case *models.Product:
		tableName = (&models.Product{}).TableName()
	case *models.Order:
		tableName = (&models.Order{}).TableName()
	}

	if tableName == "" {
		return fmt.Errorf("invalid model type")
	}

	if err := u.Db.Table(tableName).Create(model).Error; err != nil {
		return fmt.Errorf("failed to create record to %s table: %v", tableName, err)
	}

	return nil
}

func (u *UniverseRepository) Update(model interface{}) error {
	tableName := ""
	correspondingModel := make(map[string]interface{})
	var id uint
	switch m := model.(type) {
	case *models.User:
		tableName = (&models.User{}).TableName()
		correspondingModel = map[string]interface{}{
			"user_name": m.UserName,
			"email":     m.Email,
			"password":  m.Password,
		}
		id = m.ID
	case *models.Product:
		tableName = (&models.Product{}).TableName()
		correspondingModel = map[string]interface{}{
			"name":        m.Name,
			"description": m.Description,
			"price":       m.Price,
		}
		id = m.ID
	case *models.Order:
		tableName = (&models.Order{}).TableName()
		correspondingModel = map[string]interface{}{
			"user_id":    m.UserId,
			"product_id": m.ProductId,
		}
		id = m.ID
	default:
		return fmt.Errorf("invalid model type")
	}

	if tableName == "" {
		return fmt.Errorf("invalid model type")
	}

	if err := u.Db.Table(tableName).Where("id = ?", id).Updates(correspondingModel).Error; err != nil {
		return fmt.Errorf("failed to update record in %s table: %v", tableName, err)
	}

	return nil
}

// func (u *UniverseRepository) Delete(model interface{})
