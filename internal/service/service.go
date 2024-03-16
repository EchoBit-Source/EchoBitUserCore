package service

import (
	"context"

	"github.com/EchoBit-Source/EchoBitUserCore/internal/domain/dto"
)

type UserService interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, user *dto.CreateUser)
}
