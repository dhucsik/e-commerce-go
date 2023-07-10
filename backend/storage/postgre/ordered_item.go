package postgre

import (
	"strconv"
	"time"

	"github.com/dhucsik/e-commerce-go/models"
	"gorm.io/gorm"
)

type OrderedItem struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`
	OrderID   uint
	Order     Order
	ProductID uint
	Product   Product
	Quantity  uint
	Price     float64
}

type OrderedItemRepository struct {
	db *gorm.DB
}

func NewOrderedItemRepository(db *gorm.DB) *OrderedItemRepository {
	return &OrderedItemRepository{db: db}
}

/*
	func (r *OrderedItemRepository) Create(ctx context.Context, orderedItem *models.OrderedItem) (string, error) {
		model, err := toPostgreOrderedItem(orderedItem)
		if err != nil {
			return "", err
		}

		result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
		return strconv.FormatUint(uint64(model.ID), 10), result.Error
	}

	func (r *OrderedItemRepository) List(ctx context.Context) ([]models.OrderedItem, error) {
		var orderedItems []OrderedItem

		err := r.db.WithContext(ctx).Preload("Order").Preload("Product").Find(&orderedItems).Error
		if err != nil {
			return nil, err
		}

		return toOrderedItemModels(orderedItems), nil
	}

	func (r *OrderedItemRepository) Get(ctx context.Context, ID string) (models.OrderedItem, error) {
		orderedItem := new(OrderedItem)
		model := models.OrderedItem{}

		err := r.db.WithContext(ctx).Preload("Order").Preload("Product").First(orderedItem).Error
		if err != nil {
			return model, err
		}

		model = toOrderedItemModel(orderedItem)
		return model, nil
	}

	func (r *OrderedItemRepository) Update(ctx context.Context, ID string, order *models.OrderedItem) error {
		id, err := strconv.ParseUint(ID, 10, 32)
		if err != nil {
			return err
		}

		model, err := toPostgreOrderedItem(order)
		if err != nil {
			return err
		}

		model.ID = uint(id)
		return r.db.WithContext(ctx).Save(&model).Error
	}

	func (r *OrderedItemRepository) Delete(ctx context.Context, ID string) error {
		return r.db.WithContext(ctx).Delete(&OrderedItem{}, ID).Error
	}

	func toPostgreOrderedItem(o *models.OrderedItem) (OrderedItem, error) {
		orderedItem := OrderedItem{}

		orderID, err := strconv.ParseUint(o.Order.ID, 10, 32)
		if err != nil {
			return orderedItem, err
		}

		productID, err := strconv.ParseUint(o.Product.ID, 10, 32)
		if err != nil {
			return orderedItem, err
		}

		orderedItem.OrderID = uint(orderID)
		orderedItem.ProductID = uint(productID)
		orderedItem.Quantity = o.Quantity
		orderedItem.Price = o.Price

		return orderedItem, nil
	}
*/
func toOrderedItemModel(o *OrderedItem) models.OrderedItem {
	return models.OrderedItem{
		ID:       strconv.FormatUint(uint64(o.ID), 10),
		Order:    toOrderModel(&o.Order),
		Product:  toProductModel(&o.Product),
		Quantity: o.Quantity,
		Price:    o.Price,
	}
}

func toOrderedItemModels(orderedItems []OrderedItem) []models.OrderedItem {
	out := make([]models.OrderedItem, len(orderedItems))

	for i, orderedItem := range orderedItems {
		out[i] = toOrderedItemModel(&orderedItem)
	}

	return out
}

func fromCartsToOrderedItems(carts []Cart, orderID uint) []OrderedItem {
	out := make([]OrderedItem, len(carts))

	for i, cart := range carts {
		out[i] = fromCartToOrderedItem(&cart, orderID)
	}

	return out
}

func fromCartToOrderedItem(cart *Cart, orderID uint) OrderedItem {
	return OrderedItem{
		OrderID:   orderID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		Price:     cart.Product.Price * float64(cart.Quantity),
	}
}
