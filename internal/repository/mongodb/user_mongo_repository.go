package mongodb

import (
	"context"

	"github.com/EchoBit-Source/EchoBitUserCore/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{collection: collection}
}

func toMongoUser(u model.UserModel) (mu mongoUser, err error) {
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
	for _, otk := range u.OneTimePreKeys {
		mu.OneTimePreKey = append(mu.OneTimePreKey, MongoOneTimePreKey(otk))
	}
	return mu, nil
}

func toModelUser(mu mongoUser) (u model.UserModel) {
	u.ID = mu.ID.Hex()
	u.Username = mu.Username
	u.PasswordHash = mu.PasswordHash
	u.PublicKey = mu.PublicKey
	u.SignedPreKey = model.SignedPreKeyModel(mu.SignedPreKey)
	u.CreatedAt = mu.CreatedAt
	u.UpdatedAt = mu.UpdatedAt

	// Convert each MongoOneTimePreKeys to a domain.OneTimePreKeys
	for _, m := range mu.OneTimePreKey {
		u.OneTimePreKeys = append(u.OneTimePreKeys, model.OneTimePreKeyModel(m))
	}
	return u
}

func (r *MongoUserRepository) Create(ctx context.Context, u model.UserModel) (model.UserModel, error) {
	mu, err := toMongoUser(u)
	if err != nil {
		return u, err
	}
	_, err = r.collection.InsertOne(ctx, mu)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *MongoUserRepository) GetByUsername(ctx context.Context, username string) (*model.UserModel, error) {
	var mu mongoUser
	err := r.collection.FindOne(ctx, mongoUser{Username: username}).Decode(&mu)
	if err != nil {
		return nil, err
	}
	u := toModelUser(mu)
	return &u, nil
}

func (r *MongoUserRepository) GetByID(ctx context.Context, id string) (*model.UserModel, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var mu mongoUser
	err = r.collection.FindOne(ctx, mongoUser{ID: objID}).Decode(&mu)
	if err != nil {
		return nil, err
	}
	u := toModelUser(mu)
	return &u, nil
}

func (r *MongoUserRepository) Update(ctx context.Context, u model.UserModel) (model.UserModel, error) {
	mu, err := toMongoUser(u)
	if err != nil {
		return u, err
	}
	_, err = r.collection.ReplaceOne(ctx, mongoUser{ID: mu.ID}, mu)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *MongoUserRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(ctx, mongoUser{ID: objID})
	if err != nil {
		return err
	}
	return nil
}
