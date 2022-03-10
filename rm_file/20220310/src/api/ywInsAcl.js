import service from '@/utils/request'

// @Tags YwInsAcl
// @Summary 创建YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwInsAcl true "创建YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywInsAcl/createYwInsAcl [post]
export const createYwInsAcl = (data) => {
  return service({
    url: '/ywInsAcl/createYwInsAcl',
    method: 'post',
    data
  })
}

// @Tags YwInsAcl
// @Summary 删除YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwInsAcl true "删除YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywInsAcl/deleteYwInsAcl [delete]
export const deleteYwInsAcl = (data) => {
  return service({
    url: '/ywInsAcl/deleteYwInsAcl',
    method: 'delete',
    data
  })
}

// @Tags YwInsAcl
// @Summary 删除YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ywInsAcl/deleteYwInsAcl [delete]
export const deleteYwInsAclByIds = (data) => {
  return service({
    url: '/ywInsAcl/deleteYwInsAclByIds',
    method: 'delete',
    data
  })
}

// @Tags YwInsAcl
// @Summary 更新YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.YwInsAcl true "更新YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ywInsAcl/updateYwInsAcl [put]
export const updateYwInsAcl = (data) => {
  return service({
    url: '/ywInsAcl/updateYwInsAcl',
    method: 'put',
    data
  })
}

// @Tags YwInsAcl
// @Summary 用id查询YwInsAcl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.YwInsAcl true "用id查询YwInsAcl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ywInsAcl/findYwInsAcl [get]
export const findYwInsAcl = (params) => {
  return service({
    url: '/ywInsAcl/findYwInsAcl',
    method: 'get',
    params
  })
}

// @Tags YwInsAcl
// @Summary 分页获取YwInsAcl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取YwInsAcl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ywInsAcl/getYwInsAclList [get]
export const getYwInsAclList = (params) => {
  return service({
    url: '/ywInsAcl/getYwInsAclList',
    method: 'get',
    params
  })
}
