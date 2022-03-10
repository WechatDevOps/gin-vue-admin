// 自动生成模板Test
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Test 结构体
// 如果含有time.Time 请自行import time包
type Test struct {
      global.GVA_MODEL
      Profile  string `json:"profile" form:"profile" gorm:"column:profile;comment:;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Gender  *int `json:"gender" form:"gender" gorm:"column:gender;comment:;"`
      PhoneNumber  string `json:"phoneNumber" form:"phoneNumber" gorm:"column:phone_number;comment:;"`
}


