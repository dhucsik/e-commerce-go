package postgre

import (
	"context"
	"strconv"
	"time"

	"github.com/dhucsik/e-commerce-go/models"
	"gorm.io/gorm"
)

type Review struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      ``
	UpdatedAt time.Time      ``
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint
	User      User
	ProductID uint
	Product   Product
	Rating    uint
	Comment   string
}

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) List(ctx context.Context, productID string) ([]models.Review, error) {
	var reviews []Review

	err := r.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Preload("Product").Where("product_id = ?", productID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return toReviewModels(reviews), nil
}

func (r *ReviewRepository) Get(ctx context.Context, ID string) (models.Review, error) {
	review := new(Review)
	model := models.Review{}

	err := r.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("Password", "UserRole", "PhoneNumber")
	}).Preload("Product").Where("id = ?", ID).First(review).Error
	if err != nil {
		return model, err
	}

	model = toReviewModel(review)
	return model, nil
}

func (r *ReviewRepository) Create(ctx context.Context, review *models.Review) (string, error) {
	model, err := toPostgreReview(review)
	if err != nil {
		return "", err
	}

	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *ReviewRepository) Update(ctx context.Context, ID string, review *models.Review) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model, err := toPostgreReview(review)
	if err != nil {
		return err
	}

	model.ID = uint(id)
	return r.db.WithContext(ctx).Save(&model).Error
}

func (r *ReviewRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&Review{}, ID).Error
}

func toReviewModels(reviews []Review) []models.Review {
	out := make([]models.Review, len(reviews))

	for i, review := range reviews {
		out[i] = toReviewModel(&review)
	}

	return out
}

func toReviewModel(r *Review) models.Review {
	return models.Review{
		ID:      strconv.FormatUint(uint64(r.ID), 10),
		User:    toUserModel(&r.User),
		Product: toProductModel(&r.Product),
		Rating:  r.Rating,
		Comment: r.Comment,
	}
}

func toPostgreReview(r *models.Review) (Review, error) {
	review := Review{}

	userID, err := strconv.ParseUint(r.User.ID, 10, 32)
	if err != nil {
		return review, err
	}

	productID, err := strconv.ParseUint(r.Product.ID, 10, 32)
	if err != nil {
		return review, err
	}

	review.UserID = uint(userID)
	review.ProductID = uint(productID)
	review.Rating = r.Rating
	review.Comment = r.Comment

	return review, nil
}
