package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type YwUserLastLoginRouter struct {
}

// InitYwUserLastLoginRouter 初始化 YwUserLastLogin 路由信息
func (s *YwUserLastLoginRouter) InitYwUserLastLoginRouter(Router *gin.RouterGroup) {
	ywUserLastLoginRouter := Router.Group("ywUserLastLogin").Use(middleware.OperationRecord())
	ywUserLastLoginRouterWithoutRecord := Router.Group("ywUserLastLogin")
	var ywUserLastLoginApi = v1.ApiGroupApp.AutoCodeApiGroup.YwUserLastLoginApi
	{
		ywUserLastLoginRouter.POST("createYwUserLastLogin", ywUserLastLoginApi.CreateYwUserLastLogin)   // 新建YwUserLastLogin
		ywUserLastLoginRouter.DELETE("deleteYwUserLastLogin", ywUserLastLoginApi.DeleteYwUserLastLogin) // 删除YwUserLastLogin
		ywUserLastLoginRouter.DELETE("deleteYwUserLastLoginByIds", ywUserLastLoginApi.DeleteYwUserLastLoginByIds) // 批量删除YwUserLastLogin
		ywUserLastLoginRouter.PUT("updateYwUserLastLogin", ywUserLastLoginApi.UpdateYwUserLastLogin)    // 更新YwUserLastLogin
	}
	{
		ywUserLastLoginRouterWithoutRecord.GET("findYwUserLastLogin", ywUserLastLoginApi.FindYwUserLastLogin)        // 根据ID获取YwUserLastLogin
		ywUserLastLoginRouterWithoutRecord.GET("getYwUserLastLoginList", ywUserLastLoginApi.GetYwUserLastLoginList)  // 获取YwUserLastLogin列表
	}
}
