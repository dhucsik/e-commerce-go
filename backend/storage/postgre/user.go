package postgre

import (
	"context"
	"strconv"
	"time"

	"github.com/dhucsik/e-commerce-go/models"
	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time      ``
	UpdatedAt   time.Time      ``
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Username    string         `gorm:"unique"`
	Email       string         `gorm:"unique"`
	Password    string         ``
	UserRole    string         ``
	PhoneNumber string         ``
	Products    []Product      `gorm:"foreignKey:SellerID"`
	Reviews     []Review
	Orders      []Order
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (string, error) {
	model := toPostgreUser(user)
	result := r.db.WithContext(ctx).Omit("deleted_at").Create(&model)
	return strconv.FormatUint(uint64(model.ID), 10), result.Error
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (models.User, error) {
	model := models.User{}

	user := new(User)
	err := r.db.WithContext(ctx).Where("username = ?", username).First(user).Error
	if err != nil {
		return model, err
	}

	model = toUserModel(user)
	return model, nil
}

func (r *UserRepository) GetByID(ctx context.Context, ID string) (models.User, error) {
	user := new(User)
	model := models.User{}

	err := r.db.WithContext(ctx).Where("id = ?", ID).First(user).Error
	if err != nil {
		return model, err
	}

	model = toUserModel(user)
	return model, nil
}

func (r *UserRepository) Update(ctx context.Context, ID string, user *models.User) error {
	id, err := strconv.ParseUint(ID, 10, 32)
	if err != nil {
		return err
	}

	model := toPostgreUser(user)
	model.ID = uint(id)
	return r.db.Save(&model).Error
}

func (r *UserRepository) UpdatePassword(ctx context.Context, ID string, password string) error {
	return r.db.Model(&User{}).Where("id = ?", ID).Update("password", password).Error
}

func (r *UserRepository) Delete(ctx context.Context, ID string) error {
	return r.db.WithContext(ctx).Delete(&User{}, ID).Error
}

func toPostgreUser(u *models.User) User {
	return User{
		Username:    u.Username,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
		UserRole:    u.UserRole,
	}
}

func toUserModel(u *User) models.User {
	return models.User{
		ID:          strconv.FormatUint(uint64(u.ID), 10),
		Username:    u.Username,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
		UserRole:    u.UserRole,
	}
}
