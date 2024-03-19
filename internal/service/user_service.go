package service

import (
	"context"

	"github.com/EchoBit-Source/EchoBitUserCore/internal/repository"
	"github.com/EchoBit-Source/EchoBitUserCore/pkg/dto"
)

type UserService interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, user *dto.CreateUserDto) (*dto.TokenDto, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
	authService    AuthService
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, user *dto.CreateUserDto) (*dto.TokenDto, error) {
	hashedPassword, err := s.authService.GenerateHash(user.Password)
	if err != nil {
		// Handle error
		return nil, err
	}

	userModel := user.ToUser(hashedPassword)
	err = s.userRepository.Create(ctx, userModel)
	if err != nil {
		// Handle error
		return nil, err
	}

	accessToken, refreshToken, err := s.authService.GenerateTokens(userModel.ID)

	return &dto.TokenDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err
}
