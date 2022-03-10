import service from '@/utils/request'

// @Tags YwUserLastLogin
// @Summary 创建YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwUserLastLogin true "创建YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywUserLastLogin/createYwUserLastLogin [post]
export const createYwUserLastLogin = (data) => {
  return service({
    url: '/ywUserLastLogin/createYwUserLastLogin',
    method: 'post',
    data
  })
}

// @Tags YwUserLastLogin
// @Summary 删除YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwUserLastLogin true "删除YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywUserLastLogin/deleteYwUserLastLogin [delete]
export const deleteYwUserLastLogin = (data) => {
  return service({
    url: '/ywUserLastLogin/deleteYwUserLastLogin',
    method: 'delete',
    data
  })
}

// @Tags YwUserLastLogin
// @Summary 删除YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywUserLastLogin/deleteYwUserLastLogin [delete]
export const deleteYwUserLastLoginByIds = (data) => {
  return service({
    url: '/ywUserLastLogin/deleteYwUserLastLoginByIds',
    method: 'delete',
    data
  })
}

// @Tags YwUserLastLogin
// @Summary 更新YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwUserLastLogin true "更新YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywUserLastLogin/updateYwUserLastLogin [put]
export const updateYwUserLastLogin = (data) => {
  return service({
    url: '/ywUserLastLogin/updateYwUserLastLogin',
    method: 'put',
    data
  })
}

// @Tags YwUserLastLogin
// @Summary 用id查询YwUserLastLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.YwUserLastLogin true "用id查询YwUserLastLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywUserLastLogin/findYwUserLastLogin [get]
export const findYwUserLastLogin = (params) => {
  return service({
    url: '/ywUserLastLogin/findYwUserLastLogin',
    method: 'get',
    params
  })
}

// @Tags YwUserLastLogin
// @Summary 分页获取YwUserLastLogin列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取YwUserLastLogin列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywUserLastLogin/getYwUserLastLoginList [get]
export const getYwUserLastLoginList = (params) => {
  return service({
    url: '/ywUserLastLogin/getYwUserLastLoginList',
    method: 'get',
    params
  })
}
