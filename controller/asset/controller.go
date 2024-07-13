package asset

import (
	"github.com/adesupraptolaia/assetfindr/service/asset"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	assetSvc asset.Service
}

func NewAssetController(assetSvc asset.Service) *Controller {
	return &Controller{
		assetSvc: assetSvc,
	}
}

func (ctrl *Controller) CreateNewAsset(c *gin.Context) {}
func (ctrl *Controller) GetAllAssets(c *gin.Context)   {}
func (ctrl *Controller) GetAssetByID(c *gin.Context)   {}
func (ctrl *Controller) UpdateAsset(c *gin.Context)    {}
func (ctrl *Controller) DeleteAsset(c *gin.Context)    {}
