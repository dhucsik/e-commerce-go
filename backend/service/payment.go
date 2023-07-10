package service

import (
	"context"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/storage"
)

type PaymentService struct {
	repo *storage.Storage
}

func NewPaymentService(repo *storage.Storage) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (s *PaymentService) Create(ctx context.Context, payment *models.Payment) (string, error) {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return "", ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return "", ErrPermissionDenied
	}
	return s.repo.Payment.Create(ctx, payment)
}

func (s *PaymentService) List(ctx context.Context) ([]models.Payment, error) {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return nil, ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return nil, ErrPermissionDenied
	}

	return s.repo.Payment.List(ctx)
}

func (s *PaymentService) Get(ctx context.Context, ID string) (models.Payment, error) {
	payment := models.Payment{}

	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return payment, ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return payment, ErrPermissionDenied
	}

	payment, err := s.repo.Payment.Get(ctx, ID)
	return payment, err
}

func (s *PaymentService) Update(ctx context.Context, ID string, payment *models.Payment) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return ErrPermissionDenied
	}

	return s.repo.Payment.Update(ctx, ID, payment)
}

func (s *PaymentService) Delete(ctx context.Context, ID string) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return ErrPermissionDenied
	}

	return s.repo.Payment.Delete(ctx, ID)
}
