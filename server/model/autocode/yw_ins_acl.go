// 自动生成模板YwInsAcl
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
  "time"
)

// YwInsAcl 结构体
// 如果含有time.Time 请自行import time包
type YwInsAcl struct {
      global.GVA_MODEL
      Ywid  *int `json:"ywid" form:"ywid" gorm:"column:ywid;comment:id;size:10;"`
      InsId  string `json:"insId" form:"insId" gorm:"column:ins_id;comment:id;size:100;"`
      InsertTime  *time.Time `json:"添加时间" form:"添加时间" gorm:"column:insert_time;comment:;"`
      UpdateTime  *time.Time `json:"更新时间" form:"更新时间" gorm:"column:update_time;comment:;"`
}


// TableName YwInsAcl 表名
func (YwInsAcl) TableName() string {
  return "yw_ins_acl"
}

