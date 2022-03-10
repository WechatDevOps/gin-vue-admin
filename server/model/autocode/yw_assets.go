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
      InsId  string `json:"insId" form:"insId" gorm:"column:ins_id;comment:ID;size:255;"`
      InsName  string `json:"insName" form:"insName" gorm:"column:ins_name;comment:;size:255;"`
      LinkIp  string `json:"linkIp" form:"linkIp" gorm:"column:link_ip;comment:ip;size:255;"`
      Port  *int `json:"port" form:"port" gorm:"column:port;comment:;size:10;"`
      User  string `json:"user" form:"user" gorm:"column:user;comment:;size:255;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;size:255;"`
      ExpiredTime  *time.Time `json:"expiredTime" form:"expiredTime" gorm:"column:expired_time;comment:;"`
      InsertTime  *time.Time `json:"insertTime" form:"insertTime" gorm:"column:insert_time;comment:;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
}


// TableName YwAssets 表名
func (YwAssets) TableName() string {
  return "yw_assets"
}

