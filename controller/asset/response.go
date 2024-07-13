package asset

type (
	AssetResponse struct {
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
