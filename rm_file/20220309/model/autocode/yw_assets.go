// 自动生成模板YwAssets
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
  "time"
)

// YwAssets 结构体
// 如果含有time.Time 请自行import time包
type YwAssets struct {
      global.GVA_MODEL
      LinkIp  string `json:"linkIp" form:"linkIp" gorm:"column:link_ip;comment:ip;size:3000;"`
      InsId  string `json:"insId" form:"insId" gorm:"column:ins_id;comment:ID;size:255;"`
      InsName  string `json:"insName" form:"insName" gorm:"column:ins_name;comment:;size:255;"`
      InsertTime  *time.Time `json:"insertTime" form:"insertTime" gorm:"column:insert_time;comment:;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
      ExpiredTime  *time.Time `json:"expiredTime" form:"expiredTime" gorm:"column:expired_time;comment:;"`
}


// TableName YwAssets 表名
func (YwAssets) TableName() string {
  return "yw_assets"
}

