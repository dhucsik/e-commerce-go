package service

import (
	"context"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/storage"
)

type CategoryService struct {
	repo *storage.Storage
}

func NewCategoryService(repo *storage.Storage) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(ctx context.Context, category *models.Category) (string, error) {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return "", ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return "", ErrPermissionDenied
	}

	return s.repo.Category.Create(ctx, category)
}

func (s *CategoryService) List(ctx context.Context) ([]models.Category, error) {
	return s.repo.Category.List(ctx)
}

func (s *CategoryService) Get(ctx context.Context, ID string) (models.Category, error) {
	return s.repo.Category.Get(ctx, ID)
}

func (s *CategoryService) Update(ctx context.Context, ID string, category *models.Category) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return ErrPermissionDenied
	}

	return s.repo.Category.Update(ctx, ID, category)
}

func (s *CategoryService) Delete(ctx context.Context, ID string) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	if user.UserRole != models.AdminRole {
		return ErrPermissionDenied
	}

	return s.repo.Category.Delete(ctx, ID)
}
