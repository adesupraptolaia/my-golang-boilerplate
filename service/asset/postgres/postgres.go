package postgres

import (
	"context"

	"github.com/adesupraptolaia/assetfindr/service/asset"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateOne(ctx context.Context, data asset.Asset) error {
	return r.db.Table("assets").WithContext(ctx).Create(&data).Error
}

func (r *Repository) GetMany(ctx context.Context) ([]asset.Asset, error) {
	return nil, nil
}

func (r *Repository) GetByID(ctx context.Context, ID string) (*asset.Asset, error) {
	return &asset.Asset{}, nil
}

func (r *Repository) UpdateByID(ctx context.Context, ID string, data asset.Asset) error {
	return nil
}

func (r *Repository) DeleteOne(ctx context.Context, ID string) error {
	return nil
}
