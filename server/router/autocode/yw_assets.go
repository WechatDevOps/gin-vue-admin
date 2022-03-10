package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type YwAssetsRouter struct {
}

// InitYwAssetsRouter 初始化 YwAssets 路由信息
func (s *YwAssetsRouter) InitYwAssetsRouter(Router *gin.RouterGroup) {
	ywAssetsRouter := Router.Group("ywAssets").Use(middleware.OperationRecord())
	ywAssetsRouterWithoutRecord := Router.Group("ywAssets")
	var ywAssetsApi = v1.ApiGroupApp.AutoCodeApiGroup.YwAssetsApi
	{
		ywAssetsRouter.POST("createYwAssets", ywAssetsApi.CreateYwAssets)   // 新建YwAssets
		ywAssetsRouter.DELETE("deleteYwAssets", ywAssetsApi.DeleteYwAssets) // 删除YwAssets
		ywAssetsRouter.DELETE("deleteYwAssetsByIds", ywAssetsApi.DeleteYwAssetsByIds) // 批量删除YwAssets
		ywAssetsRouter.PUT("updateYwAssets", ywAssetsApi.UpdateYwAssets)    // 更新YwAssets
	}
	{
		ywAssetsRouterWithoutRecord.GET("findYwAssets", ywAssetsApi.FindYwAssets)        // 根据ID获取YwAssets
		ywAssetsRouterWithoutRecord.GET("getYwAssetsList", ywAssetsApi.GetYwAssetsList)  // 获取YwAssets列表
	}
}
