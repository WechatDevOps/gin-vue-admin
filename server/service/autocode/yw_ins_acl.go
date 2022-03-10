package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type YwInsAclService struct {
}

// CreateYwInsAcl 创建YwInsAcl记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywInsAclService *YwInsAclService) CreateYwInsAcl(ywInsAcl autocode.YwInsAcl) (err error) {
	err = global.GVA_DB.Create(&ywInsAcl).Error
	return err
}

// DeleteYwInsAcl 删除YwInsAcl记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywInsAclService *YwInsAclService)DeleteYwInsAcl(ywInsAcl autocode.YwInsAcl) (err error) {
	err = global.GVA_DB.Delete(&ywInsAcl).Error
	return err
}

// DeleteYwInsAclByIds 批量删除YwInsAcl记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywInsAclService *YwInsAclService)DeleteYwInsAclByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.YwInsAcl{},"id in ?",ids.Ids).Error
	return err
}

// UpdateYwInsAcl 更新YwInsAcl记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywInsAclService *YwInsAclService)UpdateYwInsAcl(ywInsAcl autocode.YwInsAcl) (err error) {
	err = global.GVA_DB.Save(&ywInsAcl).Error
	return err
}

// GetYwInsAcl 根据id获取YwInsAcl记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywInsAclService *YwInsAclService)GetYwInsAcl(id uint) (err error, ywInsAcl autocode.YwInsAcl) {
	err = global.GVA_DB.Where("id = ?", id).First(&ywInsAcl).Error
	return
}

// GetYwInsAclInfoList 分页获取YwInsAcl记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywInsAclService *YwInsAclService)GetYwInsAclInfoList(info autoCodeReq.YwInsAclSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.YwInsAcl{})
    var ywInsAcls []autocode.YwInsAcl
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Ywid != nil {
        db = db.Where("ywid = ?",info.Ywid)
    }
    if info.InsId != "" {
        db = db.Where("ins_id = ?",info.InsId)
    }
    if info.InsertTime != nil {
         db = db.Where("insert_time = ?",info.InsertTime)
    }
    if info.UpdateTime != nil {
         db = db.Where("update_time = ?",info.UpdateTime)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&ywInsAcls).Error
	return err, ywInsAcls, total
}
