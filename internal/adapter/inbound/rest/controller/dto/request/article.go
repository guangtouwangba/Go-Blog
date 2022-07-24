package request

type ArticleRequest struct {
	State    int    `json:"state"`
	Keyword  string `json:"keyword"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Archive  int    `json:"article"` // 是否是归档文章 0：不是 1：是 ，是归档文章则返回所有文章
}
