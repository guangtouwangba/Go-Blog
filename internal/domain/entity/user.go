package entity

type User struct {
	Email    string `json:"email,omitempty" gorm:"type:varchar(100);unique_index"`
	UserName string `json:"userName" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	PhoneNum string `json:"phoneNum,omitempty" gorm:"type:varchar(100)"`
	Describe string `json:"describe,omitempty" gorm:"type:varchar(100)"`
	Type     int    `json:"type,omitempty" gorm:"type:int(11)"`
}
