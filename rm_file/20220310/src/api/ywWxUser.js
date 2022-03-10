import service from '@/utils/request'

// @Tags YwWxUser
// @Summary 创建YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwWxUser true "创建YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywWxUser/createYwWxUser [post]
export const createYwWxUser = (data) => {
  return service({
    url: '/ywWxUser/createYwWxUser',
    method: 'post',
    data
  })
}

// @Tags YwWxUser
// @Summary 删除YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwWxUser true "删除YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywWxUser/deleteYwWxUser [delete]
export const deleteYwWxUser = (data) => {
  return service({
    url: '/ywWxUser/deleteYwWxUser',
    method: 'delete',
    data
  })
}

// @Tags YwWxUser
// @Summary 删除YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywWxUser/deleteYwWxUser [delete]
export const deleteYwWxUserByIds = (data) => {
  return service({
    url: '/ywWxUser/deleteYwWxUserByIds',
    method: 'delete',
    data
  })
}

// @Tags YwWxUser
// @Summary 更新YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwWxUser true "更新YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywWxUser/updateYwWxUser [put]
export const updateYwWxUser = (data) => {
  return service({
    url: '/ywWxUser/updateYwWxUser',
    method: 'put',
    data
  })
}

// @Tags YwWxUser
// @Summary 用id查询YwWxUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.YwWxUser true "用id查询YwWxUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywWxUser/findYwWxUser [get]
export const findYwWxUser = (params) => {
  return service({
    url: '/ywWxUser/findYwWxUser',
    method: 'get',
    params
  })
}

// @Tags YwWxUser
// @Summary 分页获取YwWxUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取YwWxUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywWxUser/getYwWxUserList [get]
export const getYwWxUserList = (params) => {
  return service({
    url: '/ywWxUser/getYwWxUserList',
    method: 'get',
    params
  })
}
