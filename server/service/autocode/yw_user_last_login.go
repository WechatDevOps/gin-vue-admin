package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type YwUserLastLoginService struct {
}

// CreateYwUserLastLogin 创建YwUserLastLogin记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywUserLastLoginService *YwUserLastLoginService) CreateYwUserLastLogin(ywUserLastLogin autocode.YwUserLastLogin) (err error) {
	err = global.GVA_DB.Create(&ywUserLastLogin).Error
	return err
}

// DeleteYwUserLastLogin 删除YwUserLastLogin记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywUserLastLoginService *YwUserLastLoginService)DeleteYwUserLastLogin(ywUserLastLogin autocode.YwUserLastLogin) (err error) {
	err = global.GVA_DB.Delete(&ywUserLastLogin).Error
	return err
}

// DeleteYwUserLastLoginByIds 批量删除YwUserLastLogin记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywUserLastLoginService *YwUserLastLoginService)DeleteYwUserLastLoginByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.YwUserLastLogin{},"id in ?",ids.Ids).Error
	return err
}

// UpdateYwUserLastLogin 更新YwUserLastLogin记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywUserLastLoginService *YwUserLastLoginService)UpdateYwUserLastLogin(ywUserLastLogin autocode.YwUserLastLogin) (err error) {
	err = global.GVA_DB.Save(&ywUserLastLogin).Error
	return err
}

// GetYwUserLastLogin 根据id获取YwUserLastLogin记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywUserLastLoginService *YwUserLastLoginService)GetYwUserLastLogin(id uint) (err error, ywUserLastLogin autocode.YwUserLastLogin) {
	err = global.GVA_DB.Where("id = ?", id).First(&ywUserLastLogin).Error
	return
}

// GetYwUserLastLoginInfoList 分页获取YwUserLastLogin记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywUserLastLoginService *YwUserLastLoginService)GetYwUserLastLoginInfoList(info autoCodeReq.YwUserLastLoginSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.YwUserLastLogin{})
    var ywUserLastLogins []autocode.YwUserLastLogin
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Ywid != nil {
        db = db.Where("ywid = ?",info.Ywid)
    }
    if info.InsId != "" {
        db = db.Where("ins_id = ?",info.InsId)
    }
    if info.LastLoginTime != nil {
         db = db.Where("last_login_time = ?",info.LastLoginTime)
    }
    if info.LastTimes != nil {
        db = db.Where("last_times = ?",info.LastTimes)
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
	err = db.Limit(limit).Offset(offset).Find(&ywUserLastLogins).Error
	return err, ywUserLastLogins, total
}
