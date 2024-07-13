package asset

import "context"

type (
	Service interface {
		CreateNewAsset(ctx context.Context)
		GetAllAssets(ctx context.Context)
		GetAssetByID(ctx context.Context)
		DeleteAsset(ctx context.Context)
	}

	service struct{}
)

func NewService() Service {
	return &service{}
}

func (s *service) CreateNewAsset(ctx context.Context) {
}
func (s *service) GetAllAssets(ctx context.Context) {
}
func (s *service) GetAssetByID(ctx context.Context) {
}
func (s *service) DeleteAsset(ctx context.Context) {
}
