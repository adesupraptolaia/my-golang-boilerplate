package asset

import (
	"time"

	"github.com/adesupraptolaia/assetfindr/service/asset"
)

type (
	CreateNewAssetRequest struct {
		Name            string  `json:"name" validate:"required"`
		Type            string  `json:"type" validate:"required"`
		Value           float64 `json:"value" validate:"required"`
		AcquisitionDate string  `json:"acquisition_date" validate:"required"`
	}
)

func (req *CreateNewAssetRequest) toAsset() (asset.Asset, error) {
	aqqDate, err := time.Parse("2006-01-02", req.AcquisitionDate)
	if err != nil {
		return asset.Asset{}, err
	}

	return asset.Asset{
		Name:            req.Name,
		Type:            req.Type,
		Value:           req.Value,
		AcquisitionDate: aqqDate,
	}, nil
}
