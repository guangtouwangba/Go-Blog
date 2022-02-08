package request

type Register struct {
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	AuthorityId  string   `json:"authorityId" gorm:"default:888"`
	AuthorityIds []string `json:"authorityIds"`
}

type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}
