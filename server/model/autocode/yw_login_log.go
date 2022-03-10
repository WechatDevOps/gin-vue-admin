// 自动生成模板YwLoginLog
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
  "time"
)

// YwLoginLog 结构体
// 如果含有time.Time 请自行import time包
type YwLoginLog struct {
      global.GVA_MODEL
      SessionId  string `json:"sessionId" form:"sessionId" gorm:"column:session_id;comment:id;size:255;"`
      Ywid  *int `json:"ywid" form:"ywid" gorm:"column:ywid;comment:id;size:10;"`
      InsId  string `json:"insId" form:"insId" gorm:"column:ins_id;comment:ID;size:255;"`
      InsName  string `json:"insName" form:"insName" gorm:"column:ins_name;comment:;size:255;"`
      YwIp  string `json:"ywIp" form:"ywIp" gorm:"column:yw_ip;comment:ip;size:255;"`
      SrcIp  string `json:"srcIp" form:"srcIp" gorm:"column:src_ip;comment:ip;size:255;"`
      SrcPort  *int `json:"srcPort" form:"srcPort" gorm:"column:src_port;comment:;size:10;"`
      InsertTime  *time.Time `json:"insertTime" form:"insertTime" gorm:"column:insert_time;comment:;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:;"`
      LoginType  *int `json:"loginType" form:"loginType" gorm:"column:login_type;comment:;size:10;"`
      ReplayUrl  string `json:"replayUrl" form:"replayUrl" gorm:"column:replay_url;comment:url;size:255;"`
}


// TableName YwLoginLog 表名
func (YwLoginLog) TableName() string {
  return "yw_login_log"
}

