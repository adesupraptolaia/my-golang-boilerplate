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
		db: db.Table("assets"),
	}
}

func (r *Repository) CreateOne(ctx context.Context, data asset.Asset) error {
	return r.db.WithContext(ctx).Create(&data).Error
}

func (r *Repository) GetMany(ctx context.Context) ([]asset.Asset, error) {
	var result []asset.Asset

	err := r.db.WithContext(ctx).Find(&result).Error

	return result, err
}

func (r *Repository) GetByID(ctx context.Context, ID string) (*asset.Asset, error) {
	var result asset.Asset

	if err := r.db.WithContext(ctx).Where("id = ?", ID).Find(&result).Error; err != nil {
		return nil, err
	}

	if result.ID == "" {
		return nil, asset.ErrAssetNotFound
	}

	return &result, nil
}

func (r *Repository) UpdateByID(ctx context.Context, ID string, data asset.Asset) error {
	return r.db.WithContext(ctx).Where("id = ?", ID).Updates(&data).Error
}

func (r *Repository) DeleteOne(ctx context.Context, ID string) error {
	data := asset.Asset{}

	return r.db.WithContext(ctx).Where("id = ?", ID).Delete(&data).Error
}
