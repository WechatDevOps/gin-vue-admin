import service from '@/utils/request'

// @Tags YwAssets
// @Summary 创建YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwAssets true "创建YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywAssets/createYwAssets [post]
export const createYwAssets = (data) => {
  return service({
    url: '/ywAssets/createYwAssets',
    method: 'post',
    data
  })
}

// @Tags YwAssets
// @Summary 删除YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwAssets true "删除YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywAssets/deleteYwAssets [delete]
export const deleteYwAssets = (data) => {
  return service({
    url: '/ywAssets/deleteYwAssets',
    method: 'delete',
    data
  })
}

// @Tags YwAssets
// @Summary 删除YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywAssets/deleteYwAssets [delete]
export const deleteYwAssetsByIds = (data) => {
  return service({
    url: '/ywAssets/deleteYwAssetsByIds',
    method: 'delete',
    data
  })
}

// @Tags YwAssets
// @Summary 更新YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwAssets true "更新YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywAssets/updateYwAssets [put]
export const updateYwAssets = (data) => {
  return service({
    url: '/ywAssets/updateYwAssets',
    method: 'put',
    data
  })
}

// @Tags YwAssets
// @Summary 用id查询YwAssets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.YwAssets true "用id查询YwAssets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywAssets/findYwAssets [get]
export const findYwAssets = (params) => {
  return service({
    url: '/ywAssets/findYwAssets',
    method: 'get',
    params
  })
}

// @Tags YwAssets
// @Summary 分页获取YwAssets列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取YwAssets列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywAssets/getYwAssetsList [get]
export const getYwAssetsList = (params) => {
  return service({
    url: '/ywAssets/getYwAssetsList',
    method: 'get',
    params
  })
}
