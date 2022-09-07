package main

import (
	"demo_cms_1/dao"
	"demo_cms_1/model"
	"demo_cms_1/routers"
)

func main() {

	//链接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	dao.DB.AutoMigrate(&model.DemoAdmin{})
	r := routers.SetupRouter()

	r.Run(":8000")
}
