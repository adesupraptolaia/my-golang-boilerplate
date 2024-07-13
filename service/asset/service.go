package asset

import "context"

type (
	Service interface {
		CreateNewAsset(ctx context.Context, data Asset) error
		GetAllAssets(ctx context.Context) ([]Asset, error)
		GetAssetByID(ctx context.Context, ID string) (*Asset, error)
		UpdateAsset(ctx context.Context, ID string, data Asset) error
		DeleteAsset(ctx context.Context, ID string) error
	}

	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateNewAsset(ctx context.Context, data Asset) error {
	return s.repo.CreateOne(ctx, data)
}

func (s *service) GetAllAssets(ctx context.Context) ([]Asset, error) {
	return nil, nil
}

func (s *service) GetAssetByID(ctx context.Context, ID string) (*Asset, error) {
	return nil, nil
}

func (s *service) UpdateAsset(ctx context.Context, ID string, data Asset) error {
	return nil
}

func (s *service) DeleteAsset(ctx context.Context, ID string) error {
	return nil
}
