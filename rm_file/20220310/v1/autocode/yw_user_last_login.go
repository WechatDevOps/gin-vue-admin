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

type YwUserLastLoginApi struct {
}

var ywUserLastLoginService = service.ServiceGroupApp.AutoCodeServiceGroup.YwUserLastLoginService


// CreateYwUserLastLogin 创建YwUserLastLogin
// @Tags YwUserLastLogin
// @Summary 创建YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwUserLastLogin true "创建YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywUserLastLogin/createYwUserLastLogin [post]
func (ywUserLastLoginApi *YwUserLastLoginApi) CreateYwUserLastLogin(c *gin.Context) {
	var ywUserLastLogin autocode.YwUserLastLogin
	_ = c.ShouldBindJSON(&ywUserLastLogin)
	if err := ywUserLastLoginService.CreateYwUserLastLogin(ywUserLastLogin); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYwUserLastLogin 删除YwUserLastLogin
// @Tags YwUserLastLogin
// @Summary 删除YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwUserLastLogin true "删除YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywUserLastLogin/deleteYwUserLastLogin [delete]
func (ywUserLastLoginApi *YwUserLastLoginApi) DeleteYwUserLastLogin(c *gin.Context) {
	var ywUserLastLogin autocode.YwUserLastLogin
	_ = c.ShouldBindJSON(&ywUserLastLogin)
	if err := ywUserLastLoginService.DeleteYwUserLastLogin(ywUserLastLogin); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYwUserLastLoginByIds 批量删除YwUserLastLogin
// @Tags YwUserLastLogin
// @Summary 批量删除YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ywUserLastLogin/deleteYwUserLastLoginByIds [delete]
func (ywUserLastLoginApi *YwUserLastLoginApi) DeleteYwUserLastLoginByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := ywUserLastLoginService.DeleteYwUserLastLoginByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYwUserLastLogin 更新YwUserLastLogin
// @Tags YwUserLastLogin
// @Summary 更新YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwUserLastLogin true "更新YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywUserLastLogin/updateYwUserLastLogin [put]
func (ywUserLastLoginApi *YwUserLastLoginApi) UpdateYwUserLastLogin(c *gin.Context) {
	var ywUserLastLogin autocode.YwUserLastLogin
	_ = c.ShouldBindJSON(&ywUserLastLogin)
	if err := ywUserLastLoginService.UpdateYwUserLastLogin(ywUserLastLogin); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYwUserLastLogin 用id查询YwUserLastLogin
// @Tags YwUserLastLogin
// @Summary 用id查询YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.YwUserLastLogin true "用id查询YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywUserLastLogin/findYwUserLastLogin [get]
func (ywUserLastLoginApi *YwUserLastLoginApi) FindYwUserLastLogin(c *gin.Context) {
	var ywUserLastLogin autocode.YwUserLastLogin
	_ = c.ShouldBindQuery(&ywUserLastLogin)
	if err, reywUserLastLogin := ywUserLastLoginService.GetYwUserLastLogin(ywUserLastLogin.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reywUserLastLogin": reywUserLastLogin}, c)
	}
}

// GetYwUserLastLoginList 分页获取YwUserLastLogin列表
// @Tags YwUserLastLogin
// @Summary 分页获取YwUserLastLogin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.YwUserLastLoginSearch true "分页获取YwUserLastLogin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywUserLastLogin/getYwUserLastLoginList [get]
func (ywUserLastLoginApi *YwUserLastLoginApi) GetYwUserLastLoginList(c *gin.Context) {
	var pageInfo autocodeReq.YwUserLastLoginSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := ywUserLastLoginService.GetYwUserLastLoginInfoList(pageInfo); err != nil {
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
