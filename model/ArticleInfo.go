package model

import "time"
//主页 显示文章概要
type ArticleInfo struct {
	Id         int64 `db:"id"`
	CategoryId int   `db:"category_id"`
	//文章摘要
	Summary      string `db:"summary"`
	Title        string
	ViewCount    uint64
	CreatTime    time.Time
	CommentCount uint64
	Username     string
}
//文章详细内容
type  ArticleDetail struct {
	ArticleInfo
	//文章详情页
	Contant string
	//分类对象
	Category
}

type  ArticleRecord struct {
	ArticleInfo
	Category
}

func NewArticleRecord(articleInfo ArticleInfo, category Category) *ArticleRecord {
	return &ArticleRecord{ArticleInfo: articleInfo, Category: category}
}