package handler

import (
	"github.com/dhucsik/e-commerce-go/service"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	UpdatePassword(c echo.Context) error
}

type ICategoryHandler interface {
	CreateCategory(c echo.Context) error
	ListCategories(c echo.Context) error
	GetCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}

type IProductHandler interface {
	CreateProduct(c echo.Context) error
	ListProducts(c echo.Context) error
	GetProduct(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type IReviewHandler interface {
	CreateReview(c echo.Context) error
	ListReviews(c echo.Context) error
	GetReview(c echo.Context) error
	UpdateReview(c echo.Context) error
	DeleteReview(c echo.Context) error
}

type IOrderHandler interface {
	MakeOrder(c echo.Context) error
	ListOrders(c echo.Context) error
	GetOrder(c echo.Context) error
	UpdateOrder(c echo.Context) error
	DeleteOrder(c echo.Context) error
}

type ICartHandler interface {
	AddToCart(c echo.Context) error
	GetUsersCart(c echo.Context) error
	DeleteFromCart(c echo.Context) error
}

type Manager struct {
	User     IUserHandler
	Category ICategoryHandler
	Product  IProductHandler
	Review   IReviewHandler
	Order    IOrderHandler
	Cart     ICartHandler
}

func NewManager(srv *service.Manager) *Manager {
	userHandler := NewUserHandler(srv)
	categoryHandler := NewCategoryHandler(srv)
	productHandler := NewProductHandler(srv)
	reviewHandler := NewReviewHandler(srv)
	orderHandler := NewOrderHandler(srv)
	cartHandler := NewCartHandler(srv)

	return &Manager{
		User:     userHandler,
		Category: categoryHandler,
		Product:  productHandler,
		Review:   reviewHandler,
		Order:    orderHandler,
		Cart:     cartHandler,
	}
}
