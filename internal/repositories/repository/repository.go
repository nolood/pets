package repository

import "context"

type Repository interface {
	Create(ctx context.Context, entity interface{}) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, entity interface{}) error
	Get(ctx context.Context, id int) (interface{}, error)
}
