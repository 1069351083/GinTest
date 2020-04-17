package service

import (
	"ginTest/dao/db"
	"ginTest/model"
)

func InsterCategoryService(category *model.Category)(categoryId int64,err error)  {
	categoryId = db.InterCategory(category)
     return
}

func QuetyAllCategoryService()(CategoryList []*model.Category,err error)  {
	CategoryList,err= db.QueryAllCategory()
	return
}
