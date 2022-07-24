package entity

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email,omitempty" gorm:"type:varchar(100);unique_index"`
	UserName string    `json:"userName" gorm:"type:varchar(100)"`
	Password string    `json:"password" gorm:"type:varchar(100)"`
	PhoneNum string    `json:"phoneNum,omitempty" gorm:"type:varchar(100)"`
	Describe string    `json:"describe,omitempty" gorm:"type:varchar(100)"`
	Type     int       `json:"type,omitempty" gorm:"type:int(11)"`
}

func (u *User) SubScribe(code string) {
	log.Printf("subScribe: %s", code)
}

func (u *User) Update() {
	log.Printf("update user")
}
