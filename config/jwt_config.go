package config

type JWTConfig struct {
	Secret string `json:"secret"`
	Expire int64  `json:"expire"`
}

func (j *JWTConfig) GetJWTSecret() string {
	return j.Secret
}

func (j *JWTConfig) GetJWTExpire() int64 {
	return j.Expire
}
