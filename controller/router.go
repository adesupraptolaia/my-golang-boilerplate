package controller

import (
	"github.com/adesupraptolaia/assetfindr/controller/asset"
	"github.com/gin-gonic/gin"
)

type (
	RouterCtrl struct {
		assetController *asset.Controller
	}
)

func RegisterRoute(router *gin.Engine, ctrl RouterCtrl) {

	router.Use(gin.Recovery())

	assetGroup := router.Group("/assets")

	assetGroup.POST("", ctrl.assetController.CreateNewAsset)
	assetGroup.GET("", ctrl.assetController.GetAllAssets)
	assetGroup.GET("/:id", ctrl.assetController.GetAssetByID)
	assetGroup.PUT("/:id", ctrl.assetController.UpdateAsset)
	assetGroup.DELETE("/:id", ctrl.assetController.DeleteAsset)
}

func NewRouterController(assetController *asset.Controller) RouterCtrl {
	return RouterCtrl{
		assetController: assetController,
	}
}
