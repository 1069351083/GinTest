package db

import (
	"fmt"
	"ginTest/model"
	"github.com/jmoiron/sqlx"
)

//添加分类
func InterCategory(category *model.Category)(CategoryId int64){
	sqlstr :="insert into category(name,no) value (?,?)"
	exec, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryNo)
	if err!=nil {
		fmt.Println(err)
		return
	}
	CategoryId,err=exec.LastInsertId()
	return

}

//获取多个分类 通过切片存储id
func  QueryCategorys(categoryIds []*int64)(category []*model.Category,err error)  {
	sqlstr, args, err := sqlx.In("select * from category where id=?", categoryIds)
	if err!=nil {
		return
	}
	err = DB.Select(category, sqlstr, args)
	if err!=nil {
		return
	}
	return

}
//获取所有分类
func QueryAllCategory()(categorys []*model.Category,err error)  {
	sqlstr :="select * from category"
	err = DB.Select(categorys, sqlstr)
	if err!=nil {
		return
	}
	return
}