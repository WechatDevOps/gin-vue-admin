package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type YwWxUserRouter struct {
}

// InitYwWxUserRouter 初始化 YwWxUser 路由信息
func (s *YwWxUserRouter) InitYwWxUserRouter(Router *gin.RouterGroup) {
	ywWxUserRouter := Router.Group("ywWxUser").Use(middleware.OperationRecord())
	ywWxUserRouterWithoutRecord := Router.Group("ywWxUser")
	var ywWxUserApi = v1.ApiGroupApp.AutoCodeApiGroup.YwWxUserApi
	{
		ywWxUserRouter.POST("createYwWxUser", ywWxUserApi.CreateYwWxUser)   // 新建YwWxUser
		ywWxUserRouter.DELETE("deleteYwWxUser", ywWxUserApi.DeleteYwWxUser) // 删除YwWxUser
		ywWxUserRouter.DELETE("deleteYwWxUserByIds", ywWxUserApi.DeleteYwWxUserByIds) // 批量删除YwWxUser
		ywWxUserRouter.PUT("updateYwWxUser", ywWxUserApi.UpdateYwWxUser)    // 更新YwWxUser
	}
	{
		ywWxUserRouterWithoutRecord.GET("findYwWxUser", ywWxUserApi.FindYwWxUser)        // 根据ID获取YwWxUser
		ywWxUserRouterWithoutRecord.GET("getYwWxUserList", ywWxUserApi.GetYwWxUserList)  // 获取YwWxUser列表
	}
}
