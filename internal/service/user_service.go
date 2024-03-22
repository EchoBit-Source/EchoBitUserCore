package service

import (
	"context"
	"errors"

	"github.com/EchoBit-Source/EchoBitUserCore/internal/repository"
	"github.com/EchoBit-Source/EchoBitUserCore/pkg/dto"
)

type UserService interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, user *dto.CreateUserDto) (*dto.TokenDto, error)
	// GetUserByUsername retrieves a user by username
	GetUserByUsername(ctx context.Context, username string) (*dto.UserDto, error)
	// GetUserByID retrieves a user by ID
	GetUserByID(ctx context.Context, id string) (*dto.UserDto, error)
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

func (s *UserServiceImpl) GetUserByUsername(ctx context.Context, username string) (*dto.UserDto, error) {
	var user *dto.UserDto
	userModel, err := s.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if userModel != nil {

		signedPreKey := dto.SignedPreKeyDto{
			Key:       userModel.SignedPreKey.Key,
			Signature: userModel.SignedPreKey.Signature,
		}

		oneTimePreKeys := make([]dto.OneTimePreKeyDto, len(userModel.OneTimePreKeys))
		for i, v := range userModel.OneTimePreKeys {
			oneTimePreKeys[i] = dto.OneTimePreKeyDto{
				Key: v.Key,
			}
		}

		user = &dto.UserDto{
			ID:             userModel.ID,
			Username:       userModel.Username,
			PublicKey:      userModel.PublicKey,
			PasswordHash:   userModel.PasswordHash,
			CreatedAt:      userModel.CreatedAt,
			UpdatedAt:      userModel.UpdatedAt,
			SignedPreKey:   signedPreKey,
			OneTimePreKeys: oneTimePreKeys,
		}
	} else {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *UserServiceImpl) GetUserByID(ctx context.Context, id string) (*dto.UserDto, error) {
	var user *dto.UserDto
	userModel, err := s.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if userModel != nil {

		signedPreKey := dto.SignedPreKeyDto{
			Key:       userModel.SignedPreKey.Key,
			Signature: userModel.SignedPreKey.Signature,
		}

		oneTimePreKeys := make([]dto.OneTimePreKeyDto, len(userModel.OneTimePreKeys))
		for i, v := range userModel.OneTimePreKeys {
			oneTimePreKeys[i] = dto.OneTimePreKeyDto{
				Key: v.Key,
			}
		}

		user = &dto.UserDto{
			ID:             userModel.ID,
			Username:       userModel.Username,
			PublicKey:      userModel.PublicKey,
			PasswordHash:   userModel.PasswordHash,
			CreatedAt:      userModel.CreatedAt,
			UpdatedAt:      userModel.UpdatedAt,
			SignedPreKey:   signedPreKey,
			OneTimePreKeys: oneTimePreKeys,
		}
	} else {
		return nil, errors.New("user not found")
	}

	return user, nil
}
