package asset

type (
	AssetResponse struct {
		Name            string  `json:"name"`
		Type            string  `json:"type"`
		Value           float64 `json:"value"`
		AcquisitionDate string  `json:"acquisition_date"`
	}
)
