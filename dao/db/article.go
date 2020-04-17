package db

import (
	"ginTest/model"
	"github.com/jmoiron/sqlx"
)

func insertArticle(detail *model.ArticleDetail)(ArticleId int64,err error)  {
	sqlstr:="insert into article(,,,,)values(?,?,?,?)"
	exec, err := DB.Exec(sqlstr, detail.CategoryName, detail.CategoryNo)
	if err!=nil {
		return
	}
	ArticleId, err = exec.LastInsertId()
	if err!=nil {
		return
	}
	return
}
//获得文章列表 进行分页
func  QueryAllArticle(pageNum,PageSize int)(ArticleList []*model.ArticleInfo)  {
	sqlstr:="select * from article order by create_time desc limit ?,?"
	err := DB.Select(ArticleList, sqlstr,pageNum,PageSize)
	if err!=nil {
		return
	}
	return

}
//查询多个文章 分页处理
func  QueryArticles(ArticleIds []int,pageNum,pageSize int)(Articles []*model.ArticleDetail)  {
	sqlstr, args, err := sqlx.In("select * from article where id=?", ArticleIds)
	if err!=nil {
		return
	}
	err = DB.Select(Articles, sqlstr, args)
	if err!=nil {
		return
	}
	return
}