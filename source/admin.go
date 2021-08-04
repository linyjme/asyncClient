package source

import (
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"niceBackend/common/global"
	"niceBackend/model"
	"time"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.User{
	{BASE_MODEL: global.BASE_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Account: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "http://qmplusimg.henrongyi.top/NICE_header.jpg", AuthorityId: "888"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.NICE_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.User{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
