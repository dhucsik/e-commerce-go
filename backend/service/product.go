package service

import (
	"context"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/storage"
)

type ProductService struct {
	repo *storage.Storage
}

func NewProductService(repo *storage.Storage) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Create(ctx context.Context, product *models.Product) (string, error) {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return "", ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole && user.UserRole != models.SellerRole {
		return "", ErrPermissionDenied
	}

	product.Seller.ID = user.UserID

	return s.repo.Product.Create(ctx, product)
}

func (s *ProductService) List(ctx context.Context, queries *models.Queries) ([]models.Product, error) {
	if queries.StartPrice == "" {
		queries.StartPrice = "0"
	}
	if queries.EndPrice == "" {
		queries.EndPrice = "10000000"
	}
	if queries.StartRating == "" {
		queries.StartRating = "0"
	}
	if queries.EndRating == "" {
		queries.EndRating = "5"
	}

	return s.repo.Product.List(ctx, queries)
}

func (s *ProductService) Update(ctx context.Context, ID string, product *models.Product) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole && user.UserRole != models.SellerRole {
		return ErrPermissionDenied
	}

	productFromDB, err := s.repo.Product.Get(ctx, ID)
	if err != nil {
		return err
	}

	if productFromDB.Seller.ID != user.UserID {
		return ErrPermissionDenied
	}

	product.Seller.ID = user.UserID

	return s.repo.Product.Update(ctx, ID, product)
}

func (s *ProductService) Get(ctx context.Context, ID string) (models.Product, error) {
	return s.repo.Product.Get(ctx, ID)
}

func (s *ProductService) Delete(ctx context.Context, ID string) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole && user.UserRole != models.SellerRole {
		return ErrPermissionDenied
	}

	productFromDB, err := s.repo.Product.Get(ctx, ID)
	if err != nil {
		return err
	}

	if productFromDB.Seller.ID != user.UserID {
		return ErrPermissionDenied
	}

	return s.repo.Product.Delete(ctx, ID)
}
