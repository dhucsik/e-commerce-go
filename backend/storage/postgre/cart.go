package postgre

import (
	"context"
	"strconv"
	"time"

	"github.com/dhucsik/e-commerce-go/models"
	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint
	User      User
	ProductID uint
	Product   Product
	Quantity  uint
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) Create(ctx context.Context, cart *models.Cart) (string, error) {
	model, err := toPostgreCart(cart)
	if err != nil {
		return "", err
	}

	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *CartRepository) List(ctx context.Context, userID string) ([]models.Cart, error) {
	var carts []Cart

	err := r.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Preload("Product").Where("user_id = ?", userID).Find(&carts).Error
	if err != nil {
		return nil, err
	}

	return toCartModels(carts), nil
}

func (r *CartRepository) Get(ctx context.Context, ID string) (models.Cart, error) {
	cart := new(Cart)
	model := models.Cart{}

	err := r.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Preload("Product").Where("id = ?", ID).First(cart).Error
	if err != nil {
		return model, err
	}

	model = toCartModel(cart)
	return model, nil
}

func (r *CartRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Cart{}, ID).Error
}

func toPostgreCart(c *models.Cart) (Cart, error) {
	cart := Cart{}

	userID, err := strconv.ParseUint(c.User.ID, 10, 32)
	if err != nil {
		return cart, err
	}

	productID, err := strconv.ParseUint(c.Product.ID, 10, 32)
	if err != nil {
		return cart, err
	}

	cart.UserID = uint(userID)
	cart.ProductID = uint(productID)
	cart.Quantity = c.Quantity

	return cart, nil
}

func toCartModels(carts []Cart) []models.Cart {
	out := make([]models.Cart, len(carts))

	for i, cart := range carts {
		out[i] = toCartModel(&cart)
	}

	return out
}

func toCartModel(c *Cart) models.Cart {
	return models.Cart{
		ID:       strconv.FormatUint(uint64(c.ID), 10),
		User:     toUserModel(&c.User),
		Product:  toProductModel(&c.Product),
		Quantity: c.Quantity,
	}
}
