package movies

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Movie, error)
}

type mongoRepository struct {
	col *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return &mongoRepository{
		col: db.Collection("movies"),
	}
}

func (r *mongoRepository) GetAll(ctx context.Context) ([]Movie, error) {
	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var movies []Movie
	if err := cursor.All(ctx, &movies); err != nil {
		return nil, err
	}
	return movies, nil
}
