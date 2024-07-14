package asset_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/adesupraptolaia/assetfindr/service/asset"
	"github.com/adesupraptolaia/assetfindr/service/asset/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	assetRepo *mocks.Repository
	service   asset.Service
}

var (
	errExpected = errors.New("expected error")
)

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) SetupTest() {
	s.assetRepo = &mocks.Repository{}
	s.service = asset.NewService(s.assetRepo)
}

func (s *ServiceTestSuite) TestCreateAsset() {
	var (
		ctx        = context.TODO()
		acqDate, _ = time.Parse("2006-01-02", "2024-07-12")
		request    = asset.Asset{
			Name:            "Macbook Air M1",
			Type:            "Laptop",
			Value:           15000000,
			AcquisitionDate: acqDate,
		}
	)

	tests := []struct {
		name          string
		req           asset.Asset
		expectedError error
		mockFunc      func()
	}{
		{
			name:          "success",
			req:           request,
			expectedError: nil,
			mockFunc: func() {
				s.assetRepo.On("CreateOne", ctx, request).Return(nil).Once()
			},
		},
		{
			name:          "failed",
			req:           request,
			expectedError: errExpected,
			mockFunc: func() {
				s.assetRepo.On("CreateOne", ctx, request).Return(errExpected).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()

			err := s.service.CreateNewAsset(ctx, tt.req)

			s.Equal(tt.expectedError, err)
		})
	}
}

func (s *ServiceTestSuite) TestGetAllAssets() {
	var (
		ctx        = context.TODO()
		acqDate, _ = time.Parse("2006-01-02", "2024-07-12")
		response   = []asset.Asset{{
			ID:              "uuid",
			Name:            "Macbook Air M1",
			Type:            "Laptop",
			Value:           15000000,
			AcquisitionDate: acqDate,
		}}
	)

	tests := []struct {
		name             string
		expectedResponse []asset.Asset
		expectedError    error
		mockFunc         func()
	}{
		{
			name:             "success",
			expectedResponse: response,
			expectedError:    nil,
			mockFunc: func() {
				s.assetRepo.On("GetMany", ctx).Return(response, nil).Once()
			},
		},
		{
			name:             "failed",
			expectedResponse: nil,
			expectedError:    errExpected,
			mockFunc: func() {
				s.assetRepo.On("GetMany", ctx).Return(nil, errExpected).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()

			resp, err := s.service.GetAllAssets(ctx)

			s.Equal(tt.expectedError, err)
			s.Equal(tt.expectedResponse, resp)
		})
	}
}

func (s *ServiceTestSuite) TestGetAssetByID() {
	var (
		ID         = "uuid"
		ctx        = context.TODO()
		acqDate, _ = time.Parse("2006-01-02", "2024-07-12")
		response   = asset.Asset{
			ID:              ID,
			Name:            "Macbook Air M1",
			Type:            "Laptop",
			Value:           15000000,
			AcquisitionDate: acqDate,
		}
	)

	tests := []struct {
		name             string
		expectedResponse *asset.Asset
		expectedError    error
		mockFunc         func()
	}{
		{
			name:             "success",
			expectedResponse: &response,
			expectedError:    nil,
			mockFunc: func() {
				s.assetRepo.On("GetByID", ctx, ID).Return(&response, nil).Once()
			},
		},
		{
			name:             "failed",
			expectedResponse: nil,
			expectedError:    errExpected,
			mockFunc: func() {
				s.assetRepo.On("GetByID", ctx, ID).Return(nil, errExpected).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()

			resp, err := s.service.GetAssetByID(ctx, ID)

			s.Equal(tt.expectedError, err)
			s.Equal(tt.expectedResponse, resp)
		})
	}
}

func (s *ServiceTestSuite) TestUpdateAsset() {
	var (
		ID         = "uuid"
		ctx        = context.TODO()
		acqDate, _ = time.Parse("2006-01-02", "2024-07-12")
		request    = asset.Asset{
			Name:            "Macbook Air M1",
			Type:            "Laptop",
			Value:           15000000,
			AcquisitionDate: acqDate,
		}
	)

	tests := []struct {
		name          string
		request       asset.Asset
		expectedError error
		mockFunc      func()
	}{
		{
			name:          "success",
			request:       request,
			expectedError: nil,
			mockFunc: func() {
				s.assetRepo.On("GetByID", ctx, ID).Return(&asset.Asset{}, nil).Once()
				s.assetRepo.On("UpdateByID", ctx, ID, request).Return(nil).Once()
			},
		},
		{
			name:          "failed when get asset by ID",
			request:       request,
			expectedError: errExpected,
			mockFunc: func() {
				s.assetRepo.On("GetByID", ctx, ID).Return(nil, errExpected).Once()
			},
		},
		{
			name:          "failed",
			request:       request,
			expectedError: errExpected,
			mockFunc: func() {
				s.assetRepo.On("GetByID", ctx, ID).Return(&asset.Asset{}, nil).Once()
				s.assetRepo.On("UpdateByID", ctx, ID, request).Return(errExpected).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()

			err := s.service.UpdateAsset(ctx, ID, tt.request)

			s.Equal(tt.expectedError, err)
		})
	}
}

func (s *ServiceTestSuite) TestDeleteAsset() {
	var (
		ID  = "uuid"
		ctx = context.TODO()
	)

	tests := []struct {
		name          string
		expectedError error
		mockFunc      func()
	}{
		{
			name:          "success",
			expectedError: nil,
			mockFunc: func() {
				s.assetRepo.On("GetByID", ctx, ID).Return(&asset.Asset{}, nil).Once()
				s.assetRepo.On("DeleteOne", ctx, ID).Return(nil).Once()
			},
		},
		{
			name:          "failed when get asset by ID",
			expectedError: errExpected,
			mockFunc: func() {
				s.assetRepo.On("GetByID", ctx, ID).Return(nil, errExpected).Once()
			},
		},
		{
			name:          "failed",
			expectedError: errExpected,
			mockFunc: func() {
				s.assetRepo.On("GetByID", ctx, ID).Return(&asset.Asset{}, nil).Once()
				s.assetRepo.On("DeleteOne", ctx, ID).Return(errExpected).Once()
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mockFunc()

			err := s.service.DeleteAsset(ctx, ID)

			s.Equal(tt.expectedError, err)
		})
	}
}
