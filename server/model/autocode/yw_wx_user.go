// 自动生成模板YwWxUser
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
  "time"
)

// YwWxUser 结构体
// 如果含有time.Time 请自行import time包
type YwWxUser struct {
      global.GVA_MODEL
      Ywid  *int `json:"ywid" form:"ywid" gorm:"column:ywid;comment:id;"`
      WxPhoto  string `json:"wxPhoto" form:"wxPhoto" gorm:"column:wx_photo;comment:;size:255;"`
      WxName  string `json:"wxName" form:"wxName" gorm:"column:wx_name;comment:;size:255;"`
      NameNotes  string `json:"nameNotes" form:"nameNotes" gorm:"column:name_notes;comment:;size:255;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;size:50;"`
      IsRegistered  *int `json:"isRegistered" form:"isRegistered" gorm:"column:is_registered;comment:;size:10;"`
      InsertTime  *time.Time `json:"insertTime" form:"insertTime" gorm:"column:insert_time;comment:;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
      LastLoginIp  string `json:"lastLoginIp" form:"lastLoginIp" gorm:"column:last_login_ip;comment:ip;size:100;"`
      LastLoginInsid  string `json:"lastLoginInsid" form:"lastLoginInsid" gorm:"column:last_login_insid;comment:ins_id;size:50;"`
      LastLoginTime  *time.Time `json:"lastLoginTime" form:"lastLoginTime" gorm:"column:last_login_time;comment:;"`
      Openid  string `json:"openid" form:"openid" gorm:"column:openid;comment:;size:255;"`
}


// TableName YwWxUser 表名
func (YwWxUser) TableName() string {
  return "yw_wx_user"
}

