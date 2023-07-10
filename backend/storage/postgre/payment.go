package postgre

import (
	"context"
	"strconv"
	"time"

	"github.com/dhucsik/e-commerce-go/models"
	"gorm.io/gorm"
)

type Payment struct {
	ID            uint           `gorm:"primaryKey"`
	CreatedAt     time.Time      ``
	UpdatedAt     time.Time      ``
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	OrderID       uint
	Order         Order
	PaymentMethod string
	TransactionID string
	Status        string
}

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) List(ctx context.Context) ([]models.Payment, error) {
	var payments []Payment

	err := r.db.WithContext(ctx).Preload("Order").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Preload("Product").Find(&payments).Error
	if err != nil {
		return nil, err
	}

	return toPaymentModels(payments), nil
}

func (r *PaymentRepository) Get(ctx context.Context, ID string) (models.Payment, error) {
	payment := new(Payment)
	model := models.Payment{}

	err := r.db.WithContext(ctx).Where("id = ?", ID).Preload("Order").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Preload("Product").First(payment).Error
	if err != nil {
		return model, err
	}

	model = toPaymentModel(payment)
	return model, nil
}

func (r *PaymentRepository) Create(ctx context.Context, payment *models.Payment) (string, error) {
	model, err := toPostgrePayment(payment)
	if err != nil {
		return "", err
	}

	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *PaymentRepository) Update(ctx context.Context, ID string, payment *models.Payment) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model, err := toPostgrePayment(payment)
	if err != nil {
		return err
	}

	model.ID = uint(id)
	return r.db.WithContext(ctx).Save(&model).Error
}

func (r *PaymentRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Payment{}, ID).Error
}

func toPostgrePayment(p *models.Payment) (Payment, error) {
	payment := Payment{}

	orderID, err := strconv.ParseUint(p.Order.ID, 10, 32)
	if err != nil {
		return payment, err
	}

	payment.OrderID = uint(orderID)
	payment.PaymentMethod = p.PaymentMethod
	payment.TransactionID = p.TransactionID
	payment.Status = p.Status

	return payment, nil
}

func toPaymentModels(payments []Payment) []models.Payment {
	out := make([]models.Payment, len(payments))

	for i, payment := range payments {
		out[i] = toPaymentModel(&payment)
	}

	return out
}

func toPaymentModel(p *Payment) models.Payment {
	return models.Payment{
		ID:            strconv.FormatUint(uint64(p.ID), 10),
		Order:         toOrderModel(&p.Order),
		PaymentMethod: p.PaymentMethod,
		TransactionID: p.TransactionID,
		Status:        p.Status,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}
