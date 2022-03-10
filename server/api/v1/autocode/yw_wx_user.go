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

type YwWxUserApi struct {
}

var ywWxUserService = service.ServiceGroupApp.AutoCodeServiceGroup.YwWxUserService


// CreateYwWxUser 创建YwWxUser
// @Tags YwWxUser
// @Summary 创建YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwWxUser true "创建YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywWxUser/createYwWxUser [post]
func (ywWxUserApi *YwWxUserApi) CreateYwWxUser(c *gin.Context) {
	var ywWxUser autocode.YwWxUser
	_ = c.ShouldBindJSON(&ywWxUser)
	if err := ywWxUserService.CreateYwWxUser(ywWxUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYwWxUser 删除YwWxUser
// @Tags YwWxUser
// @Summary 删除YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwWxUser true "删除YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywWxUser/deleteYwWxUser [delete]
func (ywWxUserApi *YwWxUserApi) DeleteYwWxUser(c *gin.Context) {
	var ywWxUser autocode.YwWxUser
	_ = c.ShouldBindJSON(&ywWxUser)
	if err := ywWxUserService.DeleteYwWxUser(ywWxUser); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYwWxUserByIds 批量删除YwWxUser
// @Tags YwWxUser
// @Summary 批量删除YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ywWxUser/deleteYwWxUserByIds [delete]
func (ywWxUserApi *YwWxUserApi) DeleteYwWxUserByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := ywWxUserService.DeleteYwWxUserByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYwWxUser 更新YwWxUser
// @Tags YwWxUser
// @Summary 更新YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwWxUser true "更新YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywWxUser/updateYwWxUser [put]
func (ywWxUserApi *YwWxUserApi) UpdateYwWxUser(c *gin.Context) {
	var ywWxUser autocode.YwWxUser
	_ = c.ShouldBindJSON(&ywWxUser)
	if err := ywWxUserService.UpdateYwWxUser(ywWxUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYwWxUser 用id查询YwWxUser
// @Tags YwWxUser
// @Summary 用id查询YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.YwWxUser true "用id查询YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywWxUser/findYwWxUser [get]
func (ywWxUserApi *YwWxUserApi) FindYwWxUser(c *gin.Context) {
	var ywWxUser autocode.YwWxUser
	_ = c.ShouldBindQuery(&ywWxUser)
	if err, reywWxUser := ywWxUserService.GetYwWxUser(ywWxUser.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reywWxUser": reywWxUser}, c)
	}
}

// GetYwWxUserList 分页获取YwWxUser列表
// @Tags YwWxUser
// @Summary 分页获取YwWxUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.YwWxUserSearch true "分页获取YwWxUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywWxUser/getYwWxUserList [get]
func (ywWxUserApi *YwWxUserApi) GetYwWxUserList(c *gin.Context) {
	var pageInfo autocodeReq.YwWxUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := ywWxUserService.GetYwWxUserInfoList(pageInfo); err != nil {
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
