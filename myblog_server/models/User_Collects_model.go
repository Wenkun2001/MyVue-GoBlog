package models

import "time"

// UserCollectsModel 自定义第三张表
// 记录用户什么时候收藏了什么文章
type UserCollectsModel struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreateAt     time.Time
}
