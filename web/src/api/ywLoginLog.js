import service from '@/utils/request'

// @Tags YwLoginLog
// @Summary 创建YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwLoginLog true "创建YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywLoginLog/createYwLoginLog [post]
export const createYwLoginLog = (data) => {
  return service({
    url: '/ywLoginLog/createYwLoginLog',
    method: 'post',
    data
  })
}

// @Tags YwLoginLog
// @Summary 删除YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwLoginLog true "删除YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywLoginLog/deleteYwLoginLog [delete]
export const deleteYwLoginLog = (data) => {
  return service({
    url: '/ywLoginLog/deleteYwLoginLog',
    method: 'delete',
    data
  })
}

// @Tags YwLoginLog
// @Summary 删除YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywLoginLog/deleteYwLoginLog [delete]
export const deleteYwLoginLogByIds = (data) => {
  return service({
    url: '/ywLoginLog/deleteYwLoginLogByIds',
    method: 'delete',
    data
  })
}

// @Tags YwLoginLog
// @Summary 更新YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwLoginLog true "更新YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywLoginLog/updateYwLoginLog [put]
export const updateYwLoginLog = (data) => {
  return service({
    url: '/ywLoginLog/updateYwLoginLog',
    method: 'put',
    data
  })
}

// @Tags YwLoginLog
// @Summary 用id查询YwLoginLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.YwLoginLog true "用id查询YwLoginLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywLoginLog/findYwLoginLog [get]
export const findYwLoginLog = (params) => {
  return service({
    url: '/ywLoginLog/findYwLoginLog',
    method: 'get',
    params
  })
}

// @Tags YwLoginLog
// @Summary 分页获取YwLoginLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取YwLoginLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywLoginLog/getYwLoginLogList [get]
export const getYwLoginLogList = (params) => {
  return service({
    url: '/ywLoginLog/getYwLoginLogList',
    method: 'get',
    params
  })
}
