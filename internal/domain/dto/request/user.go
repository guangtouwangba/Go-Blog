package request

type UserRegisterRequest struct {
	Email    string `json:"email" gorm:"type:varchar(100)"`
	UserName string `json:"name" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	PhoneNum string `json:"phone,omitempty" gorm:"type:varchar(100)"`
	Describe string `json:"introduce,omitempty" gorm:"type:varchar(100)"`
	Type     int    `json:"type" gorm:"type:int(11)"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`    // 用户名
	Password string `json:"password"` // 密码
}

type AdminLoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
