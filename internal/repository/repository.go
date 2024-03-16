package repository

import (
	"context"

	"github.com/EchoBit-Source/EchoBitUserCore/internal/domain"
)

type UserRepository interface {
	// Create creates a new user in the database
	Create(ctx context.Context, user *domain.User) error
	// GetByUsername retrieves a user from the database by username
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	// GetByID retrieves a user from the database by ID
	GetByID(ctx context.Context, id string) (*domain.User, error)
	// Update updates a user in the database
	Update(ctx context.Context, user *domain.User) error
	// Delete deletes a user from the database
	Delete(ctx context.Context, id string) error
}
