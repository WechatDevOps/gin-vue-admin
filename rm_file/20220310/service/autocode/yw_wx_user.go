package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type YwWxUserService struct {
}

// CreateYwWxUser 创建YwWxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywWxUserService *YwWxUserService) CreateYwWxUser(ywWxUser autocode.YwWxUser) (err error) {
	err = global.GVA_DB.Create(&ywWxUser).Error
	return err
}

// DeleteYwWxUser 删除YwWxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywWxUserService *YwWxUserService)DeleteYwWxUser(ywWxUser autocode.YwWxUser) (err error) {
	err = global.GVA_DB.Delete(&ywWxUser).Error
	return err
}

// DeleteYwWxUserByIds 批量删除YwWxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywWxUserService *YwWxUserService)DeleteYwWxUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.YwWxUser{},"id in ?",ids.Ids).Error
	return err
}

// UpdateYwWxUser 更新YwWxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywWxUserService *YwWxUserService)UpdateYwWxUser(ywWxUser autocode.YwWxUser) (err error) {
	err = global.GVA_DB.Save(&ywWxUser).Error
	return err
}

// GetYwWxUser 根据id获取YwWxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywWxUserService *YwWxUserService)GetYwWxUser(id uint) (err error, ywWxUser autocode.YwWxUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&ywWxUser).Error
	return
}

// GetYwWxUserInfoList 分页获取YwWxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywWxUserService *YwWxUserService)GetYwWxUserInfoList(info autoCodeReq.YwWxUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.YwWxUser{})
    var ywWxUsers []autocode.YwWxUser
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&ywWxUsers).Error
	return err, ywWxUsers, total
}
