package asset_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adesupraptolaia/assetfindr/controller/asset"
	assetSvc "github.com/adesupraptolaia/assetfindr/service/asset"
	"github.com/adesupraptolaia/assetfindr/service/asset/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ControllerTestSuite struct {
	suite.Suite
	assetService *mocks.Service
	ctrl         *asset.Controller
	router       *gin.Engine
}

var (
	errExpected = errors.New("expected error")
)

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}

func (s *ControllerTestSuite) SetupTest() {
	s.assetService = &mocks.Service{}
	s.router = gin.Default()
	s.ctrl = asset.NewAssetController(s.assetService)
}

func (s *ControllerTestSuite) TestCreateAsset() {
	var req = asset.CreateNewAssetRequest{
		Name:            "Macbook Air M1",
		Type:            "Laptop",
		Value:           15000000,
		AcquisitionDate: "2024-07-12",
	}

	var reqInvalidDate = req
	reqInvalidDate.AcquisitionDate = "abc"

	s.router.POST("/assets", s.ctrl.CreateNewAsset)

	tests := []struct {
		name               string
		req                asset.CreateNewAssetRequest
		expectedStatusCode int
		expectedResponse   string
		mockFunc           func(ctx context.Context)
	}{
		{
			name:               "missing required request",
			req:                asset.CreateNewAssetRequest{},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"status":"error","data":null,"error_message":"Key: 'CreateNewAssetRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'CreateNewAssetRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag\nKey: 'CreateNewAssetRequest.Value' Error:Field validation for 'Value' failed on the 'required' tag\nKey: 'CreateNewAssetRequest.AcquisitionDate' Error:Field validation for 'AcquisitionDate' failed on the 'required' tag"}`,
			mockFunc:           func(ctx context.Context) {},
		},
		{
			name:               "invalid acquisition_date",
			req:                reqInvalidDate,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"status":"error","data":null,"error_message":"parsing time \"abc\" as \"2006-01-02\": cannot parse \"abc\" as \"2006\""}`,
			mockFunc:           func(ctx context.Context) {},
		},
		{
			name:               "failed when create",
			req:                req,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"status":"error","data":null,"error_message":"expected error"}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("CreateNewAsset", mock.Anything, mock.Anything).Return(errExpected).Once()
			},
		},
		{
			name:               "success",
			req:                req,
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   `{"status":"success","data":null,"error_message":""}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("CreateNewAsset", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc(context.Background())

			reqJSON, _ := json.Marshal(tt.req)

			req, _ := http.NewRequest(http.MethodPost, "/assets", bytes.NewBuffer(reqJSON))
			req.Header.Add("content-type", "application/json")

			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, req)

			s.Equal(tt.expectedStatusCode, w.Code)
			s.Equal(tt.expectedResponse, w.Body.String())
		})
	}
}

func (s *ControllerTestSuite) TestGetAllAssets() {
	var (
		acqDate, _ = time.Parse("2006-01-02", "2024-07-12")
		response   = []assetSvc.Asset{{
			ID:              "uuid",
			Name:            "Macbook Air M1",
			Type:            "Laptop",
			Value:           15000000,
			AcquisitionDate: acqDate,
		}}
	)

	s.router.GET("/assets", s.ctrl.GetAllAssets)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedResponse   string
		mockFunc           func(ctx context.Context)
	}{
		{
			name:               "failed",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"status":"error","data":null,"error_message":"expected error"}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("GetAllAssets", mock.Anything, mock.Anything).Return(nil, errExpected).Once()
			},
		},
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"status":"success","data":[{"id":"uuid","name":"Macbook Air M1","type":"Laptop","value":15000000,"acquisition_date":"2024-07-12"}],"error_message":""}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("GetAllAssets", mock.Anything, mock.Anything).Return(response, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc(context.Background())

			req, _ := http.NewRequest(http.MethodGet, "/assets", nil)

			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, req)

			s.Equal(tt.expectedStatusCode, w.Code)
			s.Equal(tt.expectedResponse, w.Body.String())
		})
	}
}

func (s *ControllerTestSuite) TestGetAssetByID() {
	var (
		acqDate, _ = time.Parse("2006-01-02", "2024-07-12")
		ID         = "uuid"
		response   = assetSvc.Asset{
			ID:              ID,
			Name:            "Macbook Air M1",
			Type:            "Laptop",
			Value:           15000000,
			AcquisitionDate: acqDate,
		}
	)

	s.router.GET("/assets/:id", s.ctrl.GetAssetByID)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedResponse   string
		mockFunc           func(ctx context.Context)
	}{
		{
			name:               "failed",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"status":"error","data":null,"error_message":"expected error"}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("GetAssetByID", mock.Anything, ID).Return(nil, errExpected).Once()
			},
		},
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"status":"success","data":{"id":"uuid","name":"Macbook Air M1","type":"Laptop","value":15000000,"acquisition_date":"2024-07-12"},"error_message":""}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("GetAssetByID", mock.Anything, ID).Return(&response, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc(context.Background())

			req, _ := http.NewRequest(http.MethodGet, "/assets/"+ID, nil)

			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, req)

			s.Equal(tt.expectedStatusCode, w.Code)
			s.Equal(tt.expectedResponse, w.Body.String())
		})
	}
}

func (s *ControllerTestSuite) TestUpdateAsset() {
	var (
		ID  = "uuid"
		req = asset.CreateNewAssetRequest{
			Name:            "Macbook Air M1",
			Type:            "Laptop",
			Value:           15000000,
			AcquisitionDate: "2024-07-12",
		}
	)

	var reqInvalidDate = req
	reqInvalidDate.AcquisitionDate = "abc"

	s.router.PUT("/assets/:id", s.ctrl.UpdateAsset)

	tests := []struct {
		name               string
		req                asset.CreateNewAssetRequest
		expectedStatusCode int
		expectedResponse   string
		mockFunc           func(ctx context.Context)
	}{
		{
			name:               "missing required request",
			req:                asset.CreateNewAssetRequest{},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"status":"error","data":null,"error_message":"Key: 'CreateNewAssetRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'CreateNewAssetRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag\nKey: 'CreateNewAssetRequest.Value' Error:Field validation for 'Value' failed on the 'required' tag\nKey: 'CreateNewAssetRequest.AcquisitionDate' Error:Field validation for 'AcquisitionDate' failed on the 'required' tag"}`,
			mockFunc:           func(ctx context.Context) {},
		},
		{
			name:               "invalid acquisition_date",
			req:                reqInvalidDate,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"status":"error","data":null,"error_message":"parsing time \"abc\" as \"2006-01-02\": cannot parse \"abc\" as \"2006\""}`,
			mockFunc:           func(ctx context.Context) {},
		},
		{
			name:               "failed when create",
			req:                req,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"status":"error","data":null,"error_message":"expected error"}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("UpdateAsset", mock.Anything, ID, mock.Anything).Return(errExpected).Once()
			},
		},
		{
			name:               "success",
			req:                req,
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   `{"status":"success","data":null,"error_message":""}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("UpdateAsset", mock.Anything, ID, mock.Anything).Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc(context.Background())

			reqJSON, _ := json.Marshal(tt.req)

			req, _ := http.NewRequest(http.MethodPut, "/assets/"+ID, bytes.NewBuffer(reqJSON))
			req.Header.Add("content-type", "application/json")

			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, req)

			s.Equal(tt.expectedStatusCode, w.Code)
			s.Equal(tt.expectedResponse, w.Body.String())
		})
	}
}

func (s *ControllerTestSuite) TestDeleteAsset() {
	var (
		ID = "uuid"
	)

	s.router.DELETE("/assets/:id", s.ctrl.DeleteAsset)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedResponse   string
		mockFunc           func(ctx context.Context)
	}{
		{
			name:               "failed",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"status":"error","data":null,"error_message":"expected error"}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("DeleteAsset", mock.Anything, ID).Return(errExpected).Once()
			},
		},
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"status":"success","data":null,"error_message":""}`,
			mockFunc: func(ctx context.Context) {
				s.assetService.On("DeleteAsset", mock.Anything, ID).Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc(context.Background())

			req, _ := http.NewRequest(http.MethodDelete, "/assets/"+ID, nil)

			w := httptest.NewRecorder()
			s.router.ServeHTTP(w, req)

			s.Equal(tt.expectedStatusCode, w.Code)
			s.Equal(tt.expectedResponse, w.Body.String())
		})
	}
}
