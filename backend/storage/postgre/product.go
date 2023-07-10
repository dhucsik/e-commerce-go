package postgre

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dhucsik/e-commerce-go/models"
	"gorm.io/gorm"
)

type Product struct {
	ID           uint           `gorm:"primaryKey"`
	CreatedAt    time.Time      ``
	UpdatedAt    time.Time      ``
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Title        string
	SellerID     uint
	Seller       User
	CategoryID   uint
	Category     Category
	Price        float64
	Description  string
	AvgRating    float64
	OrderedItems []OrderedItem
	Reviews      []Review
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) List(ctx context.Context, queries *models.Queries) ([]models.Product, error) {
	var products []Product

	err := r.db.WithContext(ctx).Preload("Seller", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Preload("Category").Where(`LOWER(title) LIKE ? 
		AND price >= ? AND price <= ?
		AND avg_rating >= ?
		AND avg_rating <= ?
		`, fmt.Sprintf("%%%s%%", queries.Name), queries.StartPrice, queries.EndPrice, queries.StartRating, queries.EndRating).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return toProductModels(products), nil
}

func (r *ProductRepository) Get(ctx context.Context, ID string) (models.Product, error) {
	product := new(Product)
	model := models.Product{}

	err := r.db.WithContext(ctx).Preload("Seller", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Preload("Category").Where("id = ?", ID).First(product).Error
	if err != nil {
		return model, err
	}

	model = toProductModel(product)
	return model, nil
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) (string, error) {
	model, err := toPostgreProduct(product)
	if err != nil {
		return "", err
	}

	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *ProductRepository) Update(ctx context.Context, ID string, product *models.Product) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model, err := toPostgreProduct(product)
	if err != nil {
		return err
	}

	model.ID = uint(id)
	return r.db.WithContext(ctx).Save(&model).Error
}

func (r *ProductRepository) UpdateAvgRating(ctx context.Context, ID string) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	err = r.db.Exec(`UPDATE products SET avg_rating = (
					SELECT AVG(rating) FROM reviews
					WHERE reviews.product_id = ?
				)
				WHERE id = ?`, id, id).Error
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Product{}, ID).Error
}

func toPostgreProduct(p *models.Product) (Product, error) {
	product := Product{}

	sellerID, err := strconv.ParseUint(p.Seller.ID, 10, 32)
	if err != nil {
		return product, err
	}

	categoryID, err := strconv.ParseUint(p.Category.ID, 10, 32)
	if err != nil {
		return product, err
	}

	product.Title = p.Title
	product.SellerID = uint(sellerID)
	product.CategoryID = uint(categoryID)
	product.Price = p.Price
	product.Description = p.Description

	return product, nil
}

func toProductModels(products []Product) []models.Product {
	out := make([]models.Product, len(products))

	for i, product := range products {
		out[i] = toProductModel(&product)
	}

	return out
}

func toProductModel(p *Product) models.Product {
	return models.Product{
		ID:          strconv.FormatUint(uint64(p.ID), 10),
		Title:       p.Title,
		Seller:      toUserModel(&p.Seller),
		Category:    toCategoryModel(&p.Category),
		Price:       p.Price,
		Description: p.Description,
		AvgRating:   p.AvgRating,
	}
}
