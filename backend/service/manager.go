package service

import (
	"context"
	"errors"

	"github.com/dhucsik/e-commerce-go/models"
	"github.com/dhucsik/e-commerce-go/rpc"
	"github.com/dhucsik/e-commerce-go/storage"
)

type IUserService interface {
	SignUp(ctx context.Context, user *models.User) (string, error)
	SignIn(ctx context.Context, auth *models.AuthUser) (*models.ContextUserData, error)
	UpdatePassword(ctx context.Context, req *models.UpdatePasswordReq) error
}

type IProductService interface {
	Create(ctx context.Context, product *models.Product) (string, error)
	List(ctx context.Context, queries *models.Queries) ([]models.Product, error)
	Get(ctx context.Context, ID string) (models.Product, error)
	Update(ctx context.Context, ID string, product *models.Product) error
	Delete(ctx context.Context, ID string) error
}

type ICategoryService interface {
	Create(ctx context.Context, category *models.Category) (string, error)
	List(ctx context.Context) ([]models.Category, error)
	Get(ctx context.Context, ID string) (models.Category, error)
	Update(ctx context.Context, ID string, category *models.Category) error
	Delete(ctx context.Context, ID string) error
}

type IReviewService interface {
	Create(ctx context.Context, review *models.Review) (string, error)
	List(ctx context.Context, productID string) ([]models.Review, error)
	Get(ctx context.Context, ID string) (models.Review, error)
	Update(ctx context.Context, ID string, review *models.Review) error
	Delete(ctx context.Context, ID string) error
}

type IOrderService interface {
	MakeOrder(ctx context.Context, order *models.Order) error
	List(ctx context.Context) ([]models.Order, error)
	Get(ctx context.Context, ID string) (models.Order, error)
	Update(ctx context.Context, ID string, order *models.Order) error
	Delete(ctx context.Context, ID string) error
}

type ICartService interface {
	AddProductToCart(ctx context.Context, cart *models.Cart) (string, error)
	GetUsersCart(ctx context.Context) ([]models.Cart, error)
	DeleteProductFromCart(ctx context.Context, ID string) error
}

type IPaymentService interface {
	Create(ctx context.Context, payment *models.Payment) (string, error)
	List(ctx context.Context) ([]models.Payment, error)
	Get(ctx context.Context, ID string) (models.Payment, error)
	Update(ctx context.Context, ID string, payment *models.Payment) error
	Delete(ctx context.Context, ID string) error
}

type IMailService interface {
	SendMail(ctx context.Context, mail *models.Mail) error
}

type Manager struct {
	User     IUserService
	Product  IProductService
	Category ICategoryService
	Review   IReviewService
	Order    IOrderService
	Cart     ICartService
	Payment  IPaymentService
	Mail     IMailService
}

func NewManager(storage *storage.Storage, mailClient *rpc.MailClient) (*Manager, error) {
	if storage == nil {
		return nil, errors.New("no storage provided")
	}

	if mailClient == nil {
		return nil, errors.New("no grpc client provided")
	}

	userSrv := NewUserService(storage)
	productSrv := NewProductService(storage)
	categorySrv := NewCategoryService(storage)
	reviewSrv := NewReviewService(storage)
	orderSrv := NewOrderService(storage)
	cartSrv := NewCartService(storage)
	paymentSrv := NewPaymentService(storage)
	mailSrv := NewMailService(mailClient)

	return &Manager{
		User:     userSrv,
		Product:  productSrv,
		Category: categorySrv,
		Review:   reviewSrv,
		Order:    orderSrv,
		Cart:     cartSrv,
		Payment:  paymentSrv,
		Mail:     mailSrv,
	}, nil
}
