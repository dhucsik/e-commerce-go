package models

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password,omitempty"`
	UserRole    string `json:"user_role,omitempty"`
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatePasswordReq struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

const AdminRole = "admin"
const SellerRole = "seller"
const UserRole = "user"
