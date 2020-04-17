package db

import (
	"ginTest/model"
	"testing"
)

func init() {
	//如果数据库有时间类型必须设置paresTime
	dns :="root:admin@tcp(localhost:3306)/nvz?parseTime=true"
	err := Init(dns)
	if err!=nil {
		panic(err)
	}

}

func  TestInterCategory(t *testing.T) {

	category := InterCategory(&model.Category{
		CategoryId:   1,
		CategoryName: "nvz",
		CategoryNo:   1,
	})
	t.Logf("catefory:%#v",category)
}
