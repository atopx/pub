package model

type UserBase struct {
	Id         uint64 // 用户ID
	Name       string // 微信昵称
	EncodeId   string // 唯一ID
	Logo       string // 头像
	CreateTime uint64 // 创建时间
	UpdateTime uint64 // 更新时间
}
