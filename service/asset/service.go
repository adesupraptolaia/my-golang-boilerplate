package asset

import (
	"context"
	"errors"
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
	if data, err := s.repo.GetByID(ctx, ID); err != nil {
		log.Println("error when get asset by ID", ID, "err: ", err.Error())
		return err
	} else if data == nil {
		log.Println("asset not found")
		return errors.New("not found")
	}

	return s.repo.UpdateByID(ctx, ID, data)
}

func (s *service) DeleteAsset(ctx context.Context, ID string) error {
	data, err := s.repo.GetByID(ctx, ID)

	if err != nil {
		log.Println("error when get asset by ID", ID, "err: ", err.Error())
		return err
	} else if data == nil {
		log.Println("asset not found")
		return errors.New("not found")
	}

	return s.repo.DeleteOne(ctx, ID)
}
