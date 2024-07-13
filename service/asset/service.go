package asset

import (
	"context"
	"log"
)

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
	return s.repo.GetMany(ctx)
}

func (s *service) GetAssetByID(ctx context.Context, ID string) (*Asset, error) {
	return s.repo.GetByID(ctx, ID)
}

func (s *service) UpdateAsset(ctx context.Context, ID string, data Asset) error {
	if _, err := s.repo.GetByID(ctx, ID); err != nil {
		log.Println("error when get asset by ID", ID, "err: ", err.Error())
		return err
	}

	return s.repo.UpdateByID(ctx, ID, data)
}

func (s *service) DeleteAsset(ctx context.Context, ID string) error {
	if _, err := s.repo.GetByID(ctx, ID); err != nil {
		log.Println("error when get asset by ID", ID, "err: ", err.Error())
		return err
	}

	return s.repo.DeleteOne(ctx, ID)
}
