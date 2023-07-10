package postgre

import (
	"context"
	"strconv"
	"time"

	"github.com/dhucsik/e-commerce-go/models"
	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string         ``
	Products  []Product
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) (string, error) {
	model := toPostgreCategory(category)

	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *CategoryRepository) List(ctx context.Context) ([]models.Category, error) {
	var categories []Category

	err := r.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return toCategoryModels(categories), nil
}

func (r *CategoryRepository) Get(ctx context.Context, ID string) (models.Category, error) {
	category := new(Category)
	model := models.Category{}

	err := r.db.WithContext(ctx).Where("id = ?", ID).First(category).Error
	if err != nil {
		return model, err
	}

	model = toCategoryModel(category)
	return model, nil
}

func (r *CategoryRepository) Update(ctx context.Context, ID string, category *models.Category) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model := toPostgreCategory(category)
	if err != nil {
		return err
	}

	model.ID = uint(id)
	return r.db.Save(&model).Error
}

func (r *CategoryRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Category{}, ID).Error
}

func toCategoryModels(categories []Category) []models.Category {
	out := make([]models.Category, len(categories))

	for i, category := range categories {
		out[i] = toCategoryModel(&category)
	}

	return out
}

func toPostgreCategory(c *models.Category) Category {
	return Category{
		Title: c.Title,
	}
}

func toCategoryModel(c *Category) models.Category {
	return models.Category{
		ID:    strconv.FormatUint(uint64(c.ID), 10),
		Title: c.Title,
	}
}
