package service

import (
	"context"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/storage"
)

type CartService struct {
	repo *storage.Storage
}

func NewCartService(repo *storage.Storage) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) AddProductToCart(ctx context.Context, cart *models.Cart) (string, error) {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return "", ErrInvalidContextUserData
	}

	cart.User.ID = user.UserID

	return s.repo.Cart.Create(ctx, cart)
}

func (s *CartService) GetUsersCart(ctx context.Context) ([]models.Cart, error) {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return nil, ErrInvalidContextUserData
	}

	return s.repo.Cart.List(ctx, user.UserID)
}

func (s *CartService) DeleteProductFromCart(ctx context.Context, ID string) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	cartFromDB, err := s.repo.Cart.Get(ctx, ID)
	if err != nil {
		return err
	}

	if user.UserID != cartFromDB.User.ID && user.UserRole != models.AdminRole {
		return ErrPermissionDenied
	}

	return s.repo.Cart.Delete(ctx, ID)
}
