package asset

import (
	"net/http"

	"github.com/adesupraptolaia/assetfindr/service/asset"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	assetSvc asset.Service
}

func NewAssetController(assetSvc asset.Service) *Controller {
	return &Controller{
		assetSvc: assetSvc,
	}
}

var validate = validator.New()

func (ctrl *Controller) CreateNewAsset(c *gin.Context) {
	var req CreateNewAssetRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	assetData, err := req.toAsset()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	if err := ctrl.assetSvc.CreateNewAsset(c.Request.Context(), assetData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error when create new asset": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (ctrl *Controller) GetAllAssets(c *gin.Context) {}
func (ctrl *Controller) GetAssetByID(c *gin.Context) {}
func (ctrl *Controller) UpdateAsset(c *gin.Context)  {}
func (ctrl *Controller) DeleteAsset(c *gin.Context)  {}
