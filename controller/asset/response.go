package asset

import "github.com/adesupraptolaia/assetfindr/service/asset"

type (
	AssetResponse struct {
		ID              string  `json:"id"`
		Name            string  `json:"name"`
		Type            string  `json:"type"`
		Value           float64 `json:"value"`
		AcquisitionDate string  `json:"acquisition_date"`
	}

	Response struct {
		Status       string `json:"status"`
		Data         any    `json:"data"`
		ErrorMessage string `json:"error_message"`
	}
)

func ResponseSuccess(data any) Response {
	return Response{
		Status: "success",
		Data:   data,
	}
}

func ResponseError(err error) Response {
	return Response{
		Status:       "error",
		ErrorMessage: err.Error(),
	}
}

func convertToAssetResponse(data asset.Asset) AssetResponse {
	return AssetResponse{
		ID:              data.ID,
		Name:            data.Name,
		Type:            data.Type,
		Value:           data.Value,
		AcquisitionDate: data.AcquisitionDate.Format("2006-01-02"), // YYYY-MM-DD
	}
}
