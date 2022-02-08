package entity

import uuid "github.com/satori/go.uuid"

type User struct {
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID" `                   // 用户UUID
	Username    string    `json:"userName" gorm:"comment:用户登录名"`                 // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                      // 用户登录密码
	AuthorityId string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
}
