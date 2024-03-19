package repository

import (
	"context"

	"github.com/EchoBit-Source/EchoBitUserCore/internal/model"
)

type UserRepository interface {
	// Create creates a new user in the database
	Create(ctx context.Context, user *model.UserModel) error
	// GetByUsername retrieves a user from the database by username
	GetByUsername(ctx context.Context, username string) (*model.UserModel, error)
	// GetByID retrieves a user from the database by ID
	GetByID(ctx context.Context, id string) (*model.UserModel, error)
	// Update updates a user in the database
	Update(ctx context.Context, user *model.UserModel) error
	// Delete deletes a user from the database
	Delete(ctx context.Context, id string) error
}
