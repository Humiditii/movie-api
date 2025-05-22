package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, dto interface{}) (interface{}, error)

	find(ctx context.Context, dto interface{}) (string, error)
}

type mongoRepository struct {
	col *mongo.Collection
}


func NewRepository(db *mongo.Database) Repository {
	return &mongoRepository{
		col: db.Collection("ff"),
	}
}

func (r *mongoRepository) Create(ctx context.Context, dto interface{}) (interface{}, error) {
	return nil, nil
}

func (r *mongoRepository) find(ctx context.Context, dto interface{}) (string, error) {
	return "", nil
}