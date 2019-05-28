package dto

// News 新闻详情 DTO
type News struct {
	Nid        int    `json:"nid"`
	Title      string `json:"title"`
	Preview    string `json:"preview,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
