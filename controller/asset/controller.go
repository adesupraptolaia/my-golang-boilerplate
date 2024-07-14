package asset

import (
	"log"
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
	var (
		req       CreateNewAssetRequest
		logPrefix = "CreateNewAsset, err: "
	)

	if err := c.BindJSON(&req); err != nil {
		log.Println(logPrefix+"error when bind request", err.Error())
		c.JSON(http.StatusBadRequest, ResponseError(err))
		return
	}

	if err := validate.Struct(req); err != nil {
		log.Println(logPrefix+"error when validate request", err.Error())
		c.JSON(http.StatusBadRequest, ResponseError(err))
		return
	}

	assetData, err := req.toAsset()
	if err != nil {
		log.Println(logPrefix+"error when convert to asset request", err.Error())
		c.JSON(http.StatusBadRequest, ResponseError(err))
		return
	}

	if err := ctrl.assetSvc.CreateNewAsset(c.Request.Context(), assetData); err != nil {
		log.Println(logPrefix+"error when create new asset", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseError(err))
		return
	}

	c.JSON(http.StatusCreated, ResponseSuccess(nil))
}

func (ctrl *Controller) GetAllAssets(c *gin.Context) {
	var (
		logPrefix = "GetAllAssets, err: "
	)

	assets, err := ctrl.assetSvc.GetAllAssets(c.Request.Context())
	if err != nil {
		log.Println(logPrefix+"error when get all assets", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseError(err))
		return
	}

	var response = make([]AssetResponse, len(assets))
	for i, val := range assets {
		response[i] = convertToAssetResponse(val)
	}

	c.JSON(http.StatusOK, ResponseSuccess(response))
}

func (ctrl *Controller) GetAssetByID(c *gin.Context) {
	var (
		logPrefix = "GetAssetByID, err: "
		ID        = c.Param("id")
	)

	asset, err := ctrl.assetSvc.GetAssetByID(c.Request.Context(), ID)
	if err != nil {
		log.Println(logPrefix+"error when get asset by id:"+ID, "err: ", err)
		c.JSON(http.StatusInternalServerError, ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(convertToAssetResponse(*asset)))
}

func (ctrl *Controller) UpdateAsset(c *gin.Context) {
	var (
		req       CreateNewAssetRequest
		logPrefix = "CreateNewAsset, err: "
		ID        = c.Param("id")
	)

	if err := c.BindJSON(&req); err != nil {
		log.Println(logPrefix+"error when bind request", err.Error())
		c.JSON(http.StatusBadRequest, ResponseError(err))
		return
	}

	if err := validate.Struct(req); err != nil {
		log.Println(logPrefix+"error when validate request", err.Error())
		c.JSON(http.StatusBadRequest, ResponseError(err))
		return
	}

	assetData, err := req.toAsset()
	if err != nil {
		log.Println(logPrefix+"error when convert to asset request", err.Error())
		c.JSON(http.StatusBadRequest, ResponseError(err))
		return
	}

	if err := ctrl.assetSvc.UpdateAsset(c.Request.Context(), ID, assetData); err != nil {
		log.Println(logPrefix+"error when create new asset", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseError(err))
		return
	}

	c.JSON(http.StatusCreated, ResponseSuccess(nil))
}

func (ctrl *Controller) DeleteAsset(c *gin.Context) {
	var (
		logPrefix = "DeleteAsset, err: "
		ID        = c.Param("id")
	)

	if err := ctrl.assetSvc.DeleteAsset(c.Request.Context(), ID); err != nil {
		log.Println(logPrefix+"error when delete asset by ID"+ID, "err:", err.Error())
		c.JSON(http.StatusInternalServerError, ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, ResponseSuccess(nil))
}
