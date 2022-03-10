package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type YwLoginLogService struct {
}

// CreateYwLoginLog 创建YwLoginLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywLoginLogService *YwLoginLogService) CreateYwLoginLog(ywLoginLog autocode.YwLoginLog) (err error) {
	err = global.GVA_DB.Create(&ywLoginLog).Error
	return err
}

// DeleteYwLoginLog 删除YwLoginLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywLoginLogService *YwLoginLogService)DeleteYwLoginLog(ywLoginLog autocode.YwLoginLog) (err error) {
	err = global.GVA_DB.Delete(&ywLoginLog).Error
	return err
}

// DeleteYwLoginLogByIds 批量删除YwLoginLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywLoginLogService *YwLoginLogService)DeleteYwLoginLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.YwLoginLog{},"id in ?",ids.Ids).Error
	return err
}

// UpdateYwLoginLog 更新YwLoginLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywLoginLogService *YwLoginLogService)UpdateYwLoginLog(ywLoginLog autocode.YwLoginLog) (err error) {
	err = global.GVA_DB.Save(&ywLoginLog).Error
	return err
}

// GetYwLoginLog 根据id获取YwLoginLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywLoginLogService *YwLoginLogService)GetYwLoginLog(id uint) (err error, ywLoginLog autocode.YwLoginLog) {
	err = global.GVA_DB.Where("id = ?", id).First(&ywLoginLog).Error
	return
}

// GetYwLoginLogInfoList 分页获取YwLoginLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywLoginLogService *YwLoginLogService)GetYwLoginLogInfoList(info autoCodeReq.YwLoginLogSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.YwLoginLog{})
    var ywLoginLogs []autocode.YwLoginLog
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&ywLoginLogs).Error
	return err, ywLoginLogs, total
}
