package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type YwAssetsService struct {
}

// CreateYwAssets 创建YwAssets记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywAssetsService *YwAssetsService) CreateYwAssets(ywAssets autocode.YwAssets) (err error) {
	err = global.GVA_DB.Create(&ywAssets).Error
	return err
}

// DeleteYwAssets 删除YwAssets记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywAssetsService *YwAssetsService)DeleteYwAssets(ywAssets autocode.YwAssets) (err error) {
	err = global.GVA_DB.Delete(&ywAssets).Error
	return err
}

// DeleteYwAssetsByIds 批量删除YwAssets记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywAssetsService *YwAssetsService)DeleteYwAssetsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.YwAssets{},"id in ?",ids.Ids).Error
	return err
}

// UpdateYwAssets 更新YwAssets记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywAssetsService *YwAssetsService)UpdateYwAssets(ywAssets autocode.YwAssets) (err error) {
	err = global.GVA_DB.Save(&ywAssets).Error
	return err
}

// GetYwAssets 根据id获取YwAssets记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywAssetsService *YwAssetsService)GetYwAssets(id uint) (err error, ywAssets autocode.YwAssets) {
	err = global.GVA_DB.Where("id = ?", id).First(&ywAssets).Error
	return
}

// GetYwAssetsInfoList 分页获取YwAssets记录
// Author [piexlmax](https://github.com/piexlmax)
func (ywAssetsService *YwAssetsService)GetYwAssetsInfoList(info autoCodeReq.YwAssetsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.YwAssets{})
    var ywAssetss []autocode.YwAssets
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.InsId != "" {
        db = db.Where("ins_id = ?",info.InsId)
    }
    if info.InsName != "" {
        db = db.Where("ins_name LIKE ?","%"+ info.InsName+"%")
    }
    if info.LinkIp != "" {
        db = db.Where("link_ip = ?",info.LinkIp)
    }
    if info.Port != nil {
        db = db.Where("port = ?",info.Port)
    }
    if info.User != "" {
        db = db.Where("user = ?",info.User)
    }
    if info.Password != "" {
        db = db.Where("password = ?",info.Password)
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
	err = db.Limit(limit).Offset(offset).Find(&ywAssetss).Error
	return err, ywAssetss, total
}
