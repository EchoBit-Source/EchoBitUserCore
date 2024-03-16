package mongodb

import (
	"context"

	"github.com/EchoBit-Source/EchoBitUserCore/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{collection: collection}
}

func toMongoUser(u domain.User) (mu mongoUser, err error) {
	mu.ID, err = primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return mu, err // Handle conversion error
	}
	mu.Username = u.Username
	mu.PasswordHash = u.PasswordHash
	mu.PublicKey = u.PublicKey
	mu.SignedPreKey = MongoSignedPreKey(u.SignedPreKey)
	mu.CreatedAt = u.CreatedAt
	mu.UpdatedAt = u.UpdatedAt

	// Convert each domain.OneTimePreKeys to a MongoOneTimePreKeys
	for _, otk := range u.OneTimePreKey {
		mu.OneTimePreKey = append(mu.OneTimePreKey, MongoOneTimePreKey(otk))
	}
	return mu, nil
}

func toDomainUser(mu mongoUser) (u domain.User) {
	u.ID = mu.ID.Hex()
	u.Username = mu.Username
	u.PasswordHash = mu.PasswordHash
	u.PublicKey = mu.PublicKey
	u.SignedPreKey = domain.SignedPreKey(mu.SignedPreKey)
	u.CreatedAt = mu.CreatedAt
	u.UpdatedAt = mu.UpdatedAt

	// Convert each MongoOneTimePreKeys to a domain.OneTimePreKeys
	for _, m := range mu.OneTimePreKey {
		u.OneTimePreKey = append(u.OneTimePreKey, domain.OneTimePreKey(m))
	}
	return u
}

func (r *MongoUserRepository) Create(ctx context.Context, user *domain.User) error {
	mu, err := toMongoUser(*user)
	if err != nil {
		return err // Handle conversion error
	}
	_, err = r.collection.InsertOne(ctx, mu)
	return err
}

func (r *MongoUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var mu mongoUser
	err := r.collection.FindOne(ctx, mongoUser{Username: username}).Decode(&mu)
	if err != nil {
		return nil, err // Handle error
	}
	domainUser := toDomainUser(mu)
	return &domainUser, nil
}
