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

type YwAssetsApi struct {
}

var ywAssetsService = service.ServiceGroupApp.AutoCodeServiceGroup.YwAssetsService


// CreateYwAssets 创建YwAssets
// @Tags YwAssets
// @Summary 创建YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwAssets true "创建YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywAssets/createYwAssets [post]
func (ywAssetsApi *YwAssetsApi) CreateYwAssets(c *gin.Context) {
	var ywAssets autocode.YwAssets
	_ = c.ShouldBindJSON(&ywAssets)
	if err := ywAssetsService.CreateYwAssets(ywAssets); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYwAssets 删除YwAssets
// @Tags YwAssets
// @Summary 删除YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwAssets true "删除YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywAssets/deleteYwAssets [delete]
func (ywAssetsApi *YwAssetsApi) DeleteYwAssets(c *gin.Context) {
	var ywAssets autocode.YwAssets
	_ = c.ShouldBindJSON(&ywAssets)
	if err := ywAssetsService.DeleteYwAssets(ywAssets); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYwAssetsByIds 批量删除YwAssets
// @Tags YwAssets
// @Summary 批量删除YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ywAssets/deleteYwAssetsByIds [delete]
func (ywAssetsApi *YwAssetsApi) DeleteYwAssetsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := ywAssetsService.DeleteYwAssetsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYwAssets 更新YwAssets
// @Tags YwAssets
// @Summary 更新YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.YwAssets true "更新YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywAssets/updateYwAssets [put]
func (ywAssetsApi *YwAssetsApi) UpdateYwAssets(c *gin.Context) {
	var ywAssets autocode.YwAssets
	_ = c.ShouldBindJSON(&ywAssets)
	if err := ywAssetsService.UpdateYwAssets(ywAssets); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYwAssets 用id查询YwAssets
// @Tags YwAssets
// @Summary 用id查询YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.YwAssets true "用id查询YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywAssets/findYwAssets [get]
func (ywAssetsApi *YwAssetsApi) FindYwAssets(c *gin.Context) {
	var ywAssets autocode.YwAssets
	_ = c.ShouldBindQuery(&ywAssets)
	if err, reywAssets := ywAssetsService.GetYwAssets(ywAssets.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reywAssets": reywAssets}, c)
	}
}

// GetYwAssetsList 分页获取YwAssets列表
// @Tags YwAssets
// @Summary 分页获取YwAssets列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.YwAssetsSearch true "分页获取YwAssets列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywAssets/getYwAssetsList [get]
func (ywAssetsApi *YwAssetsApi) GetYwAssetsList(c *gin.Context) {
	var pageInfo autocodeReq.YwAssetsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := ywAssetsService.GetYwAssetsInfoList(pageInfo); err != nil {
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
