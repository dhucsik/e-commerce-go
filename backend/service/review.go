package service

import (
	"context"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/storage"
)

type ReviewService struct {
	repo *storage.Storage
}

func NewReviewService(repo *storage.Storage) *ReviewService {
	return &ReviewService{
		repo: repo,
	}
}

func (s *ReviewService) Create(ctx context.Context, review *models.Review) (string, error) {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return "", ErrInvalidContextUserData
	}

	review.User.ID = user.UserID

	id, err := s.repo.Review.Create(ctx, review)
	if err != nil {
		return id, err
	}

	if err := s.repo.Product.UpdateAvgRating(ctx, review.Product.ID); err != nil {
		return id, err
	}

	return id, nil
}

func (s *ReviewService) List(ctx context.Context, productID string) ([]models.Review, error) {
	return s.repo.Review.List(ctx, productID)
}

func (s *ReviewService) Get(ctx context.Context, ID string) (models.Review, error) {
	return s.repo.Review.Get(ctx, ID)
}

func (s *ReviewService) Update(ctx context.Context, ID string, review *models.Review) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	reviewFromDB, err := s.repo.Review.Get(ctx, ID)
	if err != nil {
		return err
	}

	if reviewFromDB.User.ID != user.UserID && user.UserRole != models.AdminRole {
		return ErrPermissionDenied
	}

	review.User.ID = user.UserID

	if err := s.repo.Review.Update(ctx, ID, review); err != nil {
		return err
	}

	if err := s.repo.Product.UpdateAvgRating(ctx, review.Product.ID); err != nil {
		return err
	}

	return nil
}

func (s *ReviewService) Delete(ctx context.Context, ID string) error {
	user, ok := ctx.Value(models.ContextKey).(*models.ContextUserData)
	if !ok {
		return ErrInvalidContextUserData
	}

	reviewFromDB, err := s.repo.Review.Get(ctx, ID)
	if err != nil {
		return err
	}

	if user.UserRole != models.AdminRole && reviewFromDB.User.ID != user.UserID {
		return ErrPermissionDenied
	}

	if err := s.repo.Review.Delete(ctx, ID); err != nil {
		return err
	}

	if err := s.repo.Product.UpdateAvgRating(ctx, reviewFromDB.ID); err != nil {
		return err
	}

	return nil
}
