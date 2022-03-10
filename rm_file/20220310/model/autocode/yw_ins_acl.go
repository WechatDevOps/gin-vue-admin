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
      InsId  string `json:"insId" form:"insId" gorm:"column:ins_id;comment:id;size:100;"`
      Ywid  *int `json:"ywid" form:"ywid" gorm:"column:ywid;comment:id;size:10;"`
      InsertTime  *time.Time `json:"insertTime" form:"insertTime" gorm:"column:insert_time;comment:;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}


// TableName YwInsAcl 表名
func (YwInsAcl) TableName() string {
  return "yw_ins_acl"
}

