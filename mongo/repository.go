package mongo

import (
	"context"

	drivermongo "go.mongodb.org/mongo-driver/mongo"
)

type Repository[T any] struct {
	collection *drivermongo.Collection
}

func NewRepository[T any](database *drivermongo.Database, collectionName string) *Repository[T] {
	return &Repository[T]{
		collection: database.Collection(collectionName),
	}
}

func (r *Repository[T]) Create(ctx context.Context, value T) (*drivermongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, value)
}

func (r *Repository[T]) FindByID(ctx context.Context, id any) (*T, error) {
	var value T

	err := r.collection.FindOne(ctx, map[string]any{"_id": id}).Decode(&value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (r *Repository[T]) FindAll(ctx context.Context, filter any) ([]T, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var values []T
	if err := cursor.All(ctx, &values); err != nil {
		return nil, err
	}

	return values, nil
}

func (r *Repository[T]) UpdateOne(ctx context.Context, filter any, update any) (*drivermongo.UpdateResult, error) {
	return r.collection.UpdateOne(ctx, filter, update)
}

func (r *Repository[T]) DeleteOne(ctx context.Context, filter any) (*drivermongo.DeleteResult, error) {
	return r.collection.DeleteOne(ctx, filter)
}
