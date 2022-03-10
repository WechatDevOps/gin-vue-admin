package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type YwLoginLogRouter struct {
}

// InitYwLoginLogRouter 初始化 YwLoginLog 路由信息
func (s *YwLoginLogRouter) InitYwLoginLogRouter(Router *gin.RouterGroup) {
	ywLoginLogRouter := Router.Group("ywLoginLog").Use(middleware.OperationRecord())
	ywLoginLogRouterWithoutRecord := Router.Group("ywLoginLog")
	var ywLoginLogApi = v1.ApiGroupApp.AutoCodeApiGroup.YwLoginLogApi
	{
		ywLoginLogRouter.POST("createYwLoginLog", ywLoginLogApi.CreateYwLoginLog)   // 新建YwLoginLog
		ywLoginLogRouter.DELETE("deleteYwLoginLog", ywLoginLogApi.DeleteYwLoginLog) // 删除YwLoginLog
		ywLoginLogRouter.DELETE("deleteYwLoginLogByIds", ywLoginLogApi.DeleteYwLoginLogByIds) // 批量删除YwLoginLog
		ywLoginLogRouter.PUT("updateYwLoginLog", ywLoginLogApi.UpdateYwLoginLog)    // 更新YwLoginLog
	}
	{
		ywLoginLogRouterWithoutRecord.GET("findYwLoginLog", ywLoginLogApi.FindYwLoginLog)        // 根据ID获取YwLoginLog
		ywLoginLogRouterWithoutRecord.GET("getYwLoginLogList", ywLoginLogApi.GetYwLoginLogList)  // 获取YwLoginLog列表
	}
}
