package service

import (
	"context"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/storage"
)

type OrderService struct {
	repo *storage.Storage
}

func NewOrderService(repo *storage.Storage) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) MakeOrder(ctx context.Context, order *models.Order) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	order.User.ID = user.UserID
	order.User.Email = user.UserEmail

	return s.repo.Order.MakeOrder(ctx, order)
}

func (s *OrderService) List(ctx context.Context) ([]models.Order, error) {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return nil, ErrInvalidContextUserData
	}

	return s.repo.Order.List(ctx, user.UserID)
}

func (s *OrderService) Get(ctx context.Context, ID string) (models.Order, error) {
	return s.repo.Order.Get(ctx, ID)
}

func (s *OrderService) Update(ctx context.Context, ID string, order *models.Order) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return ErrPermissionDenied
	}

	return s.repo.Order.Update(ctx, ID, order)
}

func (s *OrderService) Delete(ctx context.Context, ID string) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return ErrPermissionDenied
	}

	return s.repo.Order.Delete(ctx, ID)
}
