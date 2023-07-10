package storage

import (
	"context"

	"github.com/dhucsik/e-commerce-go/config"
	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/storage/postgre"
)

type IUserRepository interface {
	Create(ctx context.Context, user *models.User) (string, error)
	Update(ctx context.Context, ID string, user *models.User) error
	UpdatePassword(ctx context.Context, ID string, password string) error
	GetByUsername(ctx context.Context, username string) (models.User, error)
	GetByID(ctx context.Context, ID string) (models.User, error)
	Delete(ctx context.Context, ID string) error
}

type IProductRepository interface {
	Create(ctx context.Context, product *models.Product) (string, error)
	List(ctx context.Context, queries *models.Queries) ([]models.Product, error)
	Get(ctx context.Context, ID string) (models.Product, error)
	Update(ctx context.Context, ID string, product *models.Product) error
	UpdateAvgRating(ctx context.Context, ID string) error
	Delete(ctx context.Context, ID string) error
}

type ICategoryRepository interface {
	Create(ctx context.Context, category *models.Category) (string, error)
	List(ctx context.Context) ([]models.Category, error)
	Get(ctx context.Context, ID string) (models.Category, error)
	Update(ctx context.Context, ID string, category *models.Category) error
	Delete(ctx context.Context, ID string) error
}

type IReviewRepository interface {
	Create(ctx context.Context, review *models.Review) (string, error)
	List(ctx context.Context, productID string) ([]models.Review, error)
	Get(ctx context.Context, ID string) (models.Review, error)
	Update(ctx context.Context, ID string, review *models.Review) error
	Delete(ctx context.Context, ID string) error
}

type IOrderRepository interface {
	MakeOrder(ctx context.Context, order *models.Order) error
	List(ctx context.Context, userID string) ([]models.Order, error)
	Get(ctx context.Context, ID string) (models.Order, error)
	Update(ctx context.Context, ID string, order *models.Order) error
	Delete(ctx context.Context, ID string) error
}

type ICartRepository interface {
	Create(ctx context.Context, cart *models.Cart) (string, error)
	List(ctx context.Context, userID string) ([]models.Cart, error)
	Get(ctx context.Context, ID string) (models.Cart, error)
	Delete(ctx context.Context, ID string) error
}

type IPaymentRepository interface {
	Create(ctx context.Context, payment *models.Payment) (string, error)
	List(ctx context.Context) ([]models.Payment, error)
	Get(ctx context.Context, ID string) (models.Payment, error)
	Update(ctx context.Context, ID string, payment *models.Payment) error
	Delete(ctx context.Context, ID string) error
}

type Storage struct {
	User     IUserRepository
	Product  IProductRepository
	Category ICategoryRepository
	Review   IReviewRepository
	Order    IOrderRepository
	Cart     ICartRepository
	Payment  IPaymentRepository
}

func New(ctx context.Context, cfg *config.Config) (*Storage, error) {
	pgDB, err := postgre.Dial(ctx, cfg.PgURL)
	if err != nil {
		return nil, err
	}

	userRepo := postgre.NewUserRepository(pgDB)
	productRepo := postgre.NewProductRepository(pgDB)
	categoryRepo := postgre.NewCategoryRepository(pgDB)
	reviewRepo := postgre.NewReviewRepository(pgDB)
	orderRepo := postgre.NewOrderRepository(pgDB)
	cartRepo := postgre.NewCartRepository(pgDB)
	paymentRepo := postgre.NewPaymentRepository(pgDB)

	storage := Storage{
		User:     userRepo,
		Product:  productRepo,
		Category: categoryRepo,
		Review:   reviewRepo,
		Order:    orderRepo,
		Cart:     cartRepo,
		Payment:  paymentRepo,
	}

	return &storage, nil
}
