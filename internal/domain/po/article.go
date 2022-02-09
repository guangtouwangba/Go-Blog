package po

import (
	"time"
)

type Article struct {
	Id          int64      `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint(20);not null;comment:'文章id'"`
	Title       string     `json:"title" gorm:"type:varchar(255);not null;comment:'文章标题'"` // 标题
	Keywords    string     `json:"keywords" gorm:"type:varchar(255);comment:'文章关键词'"`      // 文章关键词
	Content     string     `json:"content" gorm:"type:varchar(255);comment:'文章内容'"`        //
	Description string     `json:"description" gorm:"type:varchar(255);comment:'文章描述'"`    // 描述
	Author      User       `json:"author" gorm:"embedded"`                                 // 文章作者
	Numbers     int64      `json:"numbers"`                                                // 文章字数
	ImageUrl    string     `json:"image_url"`
	Category    int64      `json:"category"`                      // 分类 1：技术 2：生活 3：其他
	State       int        `json:"state"`                         // 发布状态 0: 草稿 1: 已发布 2： 删除
	Comments    []*Comment `json:"comments" gorm:"foreignkey:Id"` // 评论
	Tags        []*Tag     `json:"tags" gorm:"foreignkey:id"`     // 标签
	//LikedUsers  []User     `json:"users" gorm:"foreignkey:UUID"`  // 点赞用户
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null;autoCreateTime;comment:'创建时间'"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime;not null;autoUpdateTime;comment:'更新时间'"`
}

type Comment struct {
	Id      int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;type:bigint(20);not null;comment:'评论id'"`
	Content string    `json:"content" gorm:"column:content;type:varchar(255);not null;comment:'评论内容'"`
	Author  User      `json:"author" gorm:"many2one;embedded"` // 评论作者
	Time    time.Time `json:"time" gorm:"column:time;type:datetime;not null;comment:'评论时间'"`
}

type Tag struct {
	Id   int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint(20);not null;comment:'标签id'"`
	Name string `json:"name"`
}
