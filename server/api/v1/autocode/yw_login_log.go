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

type YwLoginLogApi struct {
}

var ywLoginLogService = service.ServiceGroupApp.AutoCodeServiceGroup.YwLoginLogService


// CreateYwLoginLog 创建YwLoginLog
// @Tags YwLoginLog
// @Summary 创建YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwLoginLog true "创建YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywLoginLog/createYwLoginLog [post]
func (ywLoginLogApi *YwLoginLogApi) CreateYwLoginLog(c *gin.Context) {
	var ywLoginLog autocode.YwLoginLog
	_ = c.ShouldBindJSON(&ywLoginLog)
	if err := ywLoginLogService.CreateYwLoginLog(ywLoginLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYwLoginLog 删除YwLoginLog
// @Tags YwLoginLog
// @Summary 删除YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwLoginLog true "删除YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywLoginLog/deleteYwLoginLog [delete]
func (ywLoginLogApi *YwLoginLogApi) DeleteYwLoginLog(c *gin.Context) {
	var ywLoginLog autocode.YwLoginLog
	_ = c.ShouldBindJSON(&ywLoginLog)
	if err := ywLoginLogService.DeleteYwLoginLog(ywLoginLog); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYwLoginLogByIds 批量删除YwLoginLog
// @Tags YwLoginLog
// @Summary 批量删除YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ywLoginLog/deleteYwLoginLogByIds [delete]
func (ywLoginLogApi *YwLoginLogApi) DeleteYwLoginLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := ywLoginLogService.DeleteYwLoginLogByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYwLoginLog 更新YwLoginLog
// @Tags YwLoginLog
// @Summary 更新YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwLoginLog true "更新YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywLoginLog/updateYwLoginLog [put]
func (ywLoginLogApi *YwLoginLogApi) UpdateYwLoginLog(c *gin.Context) {
	var ywLoginLog autocode.YwLoginLog
	_ = c.ShouldBindJSON(&ywLoginLog)
	if err := ywLoginLogService.UpdateYwLoginLog(ywLoginLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYwLoginLog 用id查询YwLoginLog
// @Tags YwLoginLog
// @Summary 用id查询YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.YwLoginLog true "用id查询YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywLoginLog/findYwLoginLog [get]
func (ywLoginLogApi *YwLoginLogApi) FindYwLoginLog(c *gin.Context) {
	var ywLoginLog autocode.YwLoginLog
	_ = c.ShouldBindQuery(&ywLoginLog)
	if err, reywLoginLog := ywLoginLogService.GetYwLoginLog(ywLoginLog.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reywLoginLog": reywLoginLog}, c)
	}
}

// GetYwLoginLogList 分页获取YwLoginLog列表
// @Tags YwLoginLog
// @Summary 分页获取YwLoginLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.YwLoginLogSearch true "分页获取YwLoginLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywLoginLog/getYwLoginLogList [get]
func (ywLoginLogApi *YwLoginLogApi) GetYwLoginLogList(c *gin.Context) {
	var pageInfo autocodeReq.YwLoginLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := ywLoginLogService.GetYwLoginLogInfoList(pageInfo); err != nil {
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
