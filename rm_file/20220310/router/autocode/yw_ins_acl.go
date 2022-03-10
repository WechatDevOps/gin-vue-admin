package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type YwInsAclRouter struct {
}

// InitYwInsAclRouter 初始化 YwInsAcl 路由信息
func (s *YwInsAclRouter) InitYwInsAclRouter(Router *gin.RouterGroup) {
	ywInsAclRouter := Router.Group("ywInsAcl").Use(middleware.OperationRecord())
	ywInsAclRouterWithoutRecord := Router.Group("ywInsAcl")
	var ywInsAclApi = v1.ApiGroupApp.AutoCodeApiGroup.YwInsAclApi
	{
		ywInsAclRouter.POST("createYwInsAcl", ywInsAclApi.CreateYwInsAcl)   // 新建YwInsAcl
		ywInsAclRouter.DELETE("deleteYwInsAcl", ywInsAclApi.DeleteYwInsAcl) // 删除YwInsAcl
		ywInsAclRouter.DELETE("deleteYwInsAclByIds", ywInsAclApi.DeleteYwInsAclByIds) // 批量删除YwInsAcl
		ywInsAclRouter.PUT("updateYwInsAcl", ywInsAclApi.UpdateYwInsAcl)    // 更新YwInsAcl
	}
	{
		ywInsAclRouterWithoutRecord.GET("findYwInsAcl", ywInsAclApi.FindYwInsAcl)        // 根据ID获取YwInsAcl
		ywInsAclRouterWithoutRecord.GET("getYwInsAclList", ywInsAclApi.GetYwInsAclList)  // 获取YwInsAcl列表
	}
}
