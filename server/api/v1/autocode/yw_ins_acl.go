package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type YwInsAclApi struct {
}

var ywInsAclService = service.ServiceGroupApp.AutoCodeServiceGroup.YwInsAclService


// CreateYwInsAcl 创建YwInsAcl
// @Tags YwInsAcl
// @Summary 创建YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwInsAcl true "创建YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywInsAcl/createYwInsAcl [post]
func (ywInsAclApi *YwInsAclApi) CreateYwInsAcl(c *gin.Context) {
	var ywInsAcl autocode.YwInsAcl
	_ = c.ShouldBindJSON(&ywInsAcl)
	if err := ywInsAclService.CreateYwInsAcl(ywInsAcl); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYwInsAcl 删除YwInsAcl
// @Tags YwInsAcl
// @Summary 删除YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwInsAcl true "删除YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywInsAcl/deleteYwInsAcl [delete]
func (ywInsAclApi *YwInsAclApi) DeleteYwInsAcl(c *gin.Context) {
	var ywInsAcl autocode.YwInsAcl
	_ = c.ShouldBindJSON(&ywInsAcl)
	if err := ywInsAclService.DeleteYwInsAcl(ywInsAcl); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYwInsAclByIds 批量删除YwInsAcl
// @Tags YwInsAcl
// @Summary 批量删除YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ywInsAcl/deleteYwInsAclByIds [delete]
func (ywInsAclApi *YwInsAclApi) DeleteYwInsAclByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := ywInsAclService.DeleteYwInsAclByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYwInsAcl 更新YwInsAcl
// @Tags YwInsAcl
// @Summary 更新YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwInsAcl true "更新YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywInsAcl/updateYwInsAcl [put]
func (ywInsAclApi *YwInsAclApi) UpdateYwInsAcl(c *gin.Context) {
	var ywInsAcl autocode.YwInsAcl
	_ = c.ShouldBindJSON(&ywInsAcl)
	if err := ywInsAclService.UpdateYwInsAcl(ywInsAcl); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYwInsAcl 用id查询YwInsAcl
// @Tags YwInsAcl
// @Summary 用id查询YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.YwInsAcl true "用id查询YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywInsAcl/findYwInsAcl [get]
func (ywInsAclApi *YwInsAclApi) FindYwInsAcl(c *gin.Context) {
	var ywInsAcl autocode.YwInsAcl
	_ = c.ShouldBindQuery(&ywInsAcl)
	if err, reywInsAcl := ywInsAclService.GetYwInsAcl(ywInsAcl.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reywInsAcl": reywInsAcl}, c)
	}
}

// GetYwInsAclList 分页获取YwInsAcl列表
// @Tags YwInsAcl
// @Summary 分页获取YwInsAcl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.YwInsAclSearch true "分页获取YwInsAcl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywInsAcl/getYwInsAclList [get]
func (ywInsAclApi *YwInsAclApi) GetYwInsAclList(c *gin.Context) {
	var pageInfo autocodeReq.YwInsAclSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := ywInsAclService.GetYwInsAclInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
