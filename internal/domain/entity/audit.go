package entity

import "time"

type Audit struct {
	CreateBy    time.Time `json:"createBy"`
	CreateTime  time.Time `json:"createTime"`
	UpdateBy    string    `json:"updateBy"`
	UpdateTime  time.Time `json:"updateTime"`
	IsDeleted   bool      `json:"isDeleted"`
	DeletedTime time.Time `json:"deletedTime"`
}
