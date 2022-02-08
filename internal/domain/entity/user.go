package entity

type User struct {
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	UserName string `json:"userName" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	PhoneNum string `json:"phoneNum" gorm:"type:varchar(100)"`
	Describe string `json:"describe" gorm:"type:varchar(100)"`
	Type     int    `json:"type" gorm:"type:int(11)"`
}
