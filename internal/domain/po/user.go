package po

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uuid     uuid.UUID `json:"uuid" gorm:"type:varchar(100);primary_key"`
	Email    string    `json:"email" gorm:"type:varchar(100);unique_index"`
	UserName string    `json:"userName" gorm:"type:varchar(100)"`
	Password string    `json:"password" gorm:"type:varchar(100)"`
	PhoneNum string    `json:"phoneNum" gorm:"type:varchar(100)"`
	Describe string    `json:"describe" gorm:"type:varchar(100)"`
	Type     int       `json:"type" gorm:"type:int(11)"`                              // 0:普通用户 1:管理员
	CreateAt time.Time `json:"createAt" gorm:"type:datetime;autoCreateTime;not null"` // 创建时间
	UpdateAt time.Time `json:"updateAt" gorm:"type:datetime;autoUpdateTime;not null"` // 更新时间
}

type UserClaims struct {
	*jwt.StandardClaims
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Uuid = uuid.NewV4()
	return
}
