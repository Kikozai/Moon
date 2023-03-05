package models 

import (
	"time"

	db "github.com/Kikozai/Moon/db/sqlc"
)


type SignInInput struct { 
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct { 
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role string `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilteredResponse(user db.User) UserResponse{
	return UserResponse{ 
		ID: user.ID.String(),
		Email: user.Email,
		Name: user.Name,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}