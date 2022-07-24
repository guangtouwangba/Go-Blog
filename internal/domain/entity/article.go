package entity

import (
	"Go-Blog/internal/adapter/inbound/rest/listener"
	uuid "github.com/satori/go.uuid"
)

type Article struct {
	Id          uuid.UUID `json:"id"`
	Code        string    `json:"code" gorm:"type:varchar(100)"`
	Title       string    `json:"title"`       // 标题
	Keywords    string    `json:"keywords"`    // 文章关键词
	Content     string    `json:"content"`     //
	Description string    `json:"description"` // 描述
	Author      string    `json:"author"`      // 文章作者
	Numbers     int64     `json:"numbers"`     // 文章字数
	ImageUrl    string    `json:"image_url"`
	Category    int64     `json:"category"`    // 分类 1：技术 2：生活 3：其他
	State       int64     `json:"state"`       // 发布状态 0: 草稿 1: 已发布 2： 删除
	Comments    []Comment `json:"comments"`    // 评论
	Tags        []Tag     `json:"tags"`        // 标签
	LikedUsers  []User    `json:"liked_users"` // 点赞用户
}

type Comment struct {
	Observers []listener.Observer `json:"observers"`
	Id        uuid.UUID           `json:"id"`
	Code      string              `json:"code" gorm:"type:varchar(100)"`
	Content   string              `json:"content"`
	Author    string              `json:"author"`
	Time      string              `json:"time"`
}

func (c *Comment) Register(observer listener.Observer) {
	c.Observers = append(c.Observers, observer)
}

func (c *Comment) Unregister(observer listener.Observer) {
	for i, o := range c.Observers {
		if o == observer {
			c.Observers = append(c.Observers[:i], c.Observers[i+1:]...)
			break
		}
	}
}

func (c *Comment) Notify() {
	for _, observer := range c.Observers {
		observer.Update()
	}
}

type Tag struct {
	Name string `json:"name"`
}
