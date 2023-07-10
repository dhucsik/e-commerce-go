package postgre

import (
	"context"
	"strconv"
	"time"

	"github.com/dhucsik/e-commerce-go/models"
	"gorm.io/gorm"
)

type Order struct {
	ID           uint           `gorm:"primaryKey"`
	CreatedAt    time.Time      ``
	UpdatedAt    time.Time      ``
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserID       uint
	User         User
	TotalPrice   float64
	Address      string
	Payments     []Payment
	OrderedItems []OrderedItem
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) List(ctx context.Context, userID string) ([]models.Order, error) {
	var orders []Order

	err := r.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return toOrderModels(orders), nil
}

func (r *OrderRepository) Get(ctx context.Context, ID string) (models.Order, error) {
	order := new(Order)
	model := models.Order{}

	err := r.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Where("id = ?", ID).First(order).Error
	if err != nil {
		return model, err
	}

	model = toOrderModel(order)
	return model, nil
}

func (r *OrderRepository) Update(ctx context.Context, ID string, order *models.Order) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model, err := toPostgreOrder(order)
	if err != nil {
		return err
	}

	model.ID = uint(id)
	return r.db.WithContext(ctx).Save(&model).Error
}

func (r *OrderRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Order{}, ID).Error
}

func (r *OrderRepository) MakeOrder(ctx context.Context, order *models.Order) error {
	model, err := toPostgreOrder(order)
	if err != nil {
		return err
	}

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var carts []Cart
	err = r.db.WithContext(ctx).Preload("User").Preload("Product").Where("user_id = ?", model.UserID).Find(&carts).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	totalPrice := totalPriceOfCarts(carts)
	model.TotalPrice = totalPrice

	err = r.db.WithContext(ctx).Omit("deleted_at").Create(&model).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	orderedItems := fromCartsToOrderedItems(carts, model.ID)

	err = r.db.WithContext(ctx).Omit("deleted_at").Create(&orderedItems).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = r.db.WithContext(ctx).Exec("DELETE FROM carts WHERE user_id = ?", model.UserID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func toOrderModels(orders []Order) []models.Order {
	out := make([]models.Order, len(orders))

	for i, order := range orders {
		out[i] = toOrderModel(&order)
	}

	return out
}

func toOrderModel(o *Order) models.Order {
	return models.Order{
		ID:         strconv.FormatUint(uint64(o.ID), 10),
		User:       toUserModel(&o.User),
		TotalPrice: o.TotalPrice,
		Date:       o.CreatedAt,
		Address:    o.Address,

		OrderedItems: toOrderedItemModels(o.OrderedItems),
	}
}

func toPostgreOrder(o *models.Order) (Order, error) {
	order := Order{}

	userID, err := strconv.ParseUint(o.User.ID, 10, 32)
	if err != nil {
		return order, err
	}

	order.UserID = uint(userID)
	order.TotalPrice = o.TotalPrice
	order.Address = o.Address

	return order, nil
}

func totalPriceOfCarts(carts []Cart) float64 {
	sum := float64(0)
	for _, cart := range carts {
		sum += (cart.Product.Price) * float64(cart.Quantity)
	}

	return sum
}
