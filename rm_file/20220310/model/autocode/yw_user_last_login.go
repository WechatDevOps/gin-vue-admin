// 自动生成模板YwUserLastLogin
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
  "time"
)

// YwUserLastLogin 结构体
// 如果含有time.Time 请自行import time包
type YwUserLastLogin struct {
      global.GVA_MODEL
      InsId  string `json:"insId" form:"insId" gorm:"column:ins_id;comment:id;size:100;"`
      InsertTime  *time.Time `json:"insertTime" form:"insertTime" gorm:"column:insert_time;comment:;"`
      LastLoginTime  *time.Time `json:"lastLoginTime" form:"lastLoginTime" gorm:"column:last_login_time;comment:;"`
      LastTimes  *int `json:"lastTimes" form:"lastTimes" gorm:"column:last_times;comment:;size:19;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
      Ywid  *int `json:"ywid" form:"ywid" gorm:"column:ywid;comment:id;size:10;"`
}


// TableName YwUserLastLogin 表名
func (YwUserLastLogin) TableName() string {
  return "yw_user_last_login"
}

