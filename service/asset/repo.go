package asset

import "context"

type (
	Repository interface {
		CreateOne(ctx context.Context, data Asset) error
		GetMany(ctx context.Context) ([]Asset, error)
		GetByID(ctx context.Context, ID string) (*Asset, error)
		UpdateByID(ctx context.Context, ID string, data Asset) error
		DeleteOne(ctx context.Context, ID string) error
	}
)
