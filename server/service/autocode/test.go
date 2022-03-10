package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type TestService struct {
}

// CreateTest 创建Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService) CreateTest(test autocode.Test) (err error) {
	err = global.GVA_DB.Create(&test).Error
	return err
}

// DeleteTest 删除Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)DeleteTest(test autocode.Test) (err error) {
	err = global.GVA_DB.Delete(&test).Error
	return err
}

// DeleteTestByIds 批量删除Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)DeleteTestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Test{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTest 更新Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)UpdateTest(test autocode.Test) (err error) {
	err = global.GVA_DB.Save(&test).Error
	return err
}

// GetTest 根据id获取Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)GetTest(id uint) (err error, test autocode.Test) {
	err = global.GVA_DB.Where("id = ?", id).First(&test).Error
	return
}

// GetTestInfoList 分页获取Test记录
// Author [piexlmax](https://github.com/piexlmax)
func (testService *TestService)GetTestInfoList(info autoCodeReq.TestSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Test{})
    var tests []autocode.Test
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Profile != "" {
        db = db.Where("profile = ?",info.Profile)
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
    if info.Gender != nil {
        db = db.Where("gender = ?",info.Gender)
    }
    if info.PhoneNumber != "" {
        db = db.Where("phone_number = ?",info.PhoneNumber)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&tests).Error
	return err, tests, total
}
