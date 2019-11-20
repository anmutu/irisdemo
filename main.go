package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"irisdemo20191120/web/controller"
)

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	//加载模板文件
	app.RegisterView(iris.HTML("./web/view", ".html"))
	// 注册控制器
	mvc.New(app).Handle(new(controller.MovieController))
	app.Run(
		//开启web服务
		iris.Addr("localhost:8080"),
		// 按下CTRL / CMD + C时跳过错误的服务器：
		iris.WithoutServerError(iris.ErrServerClosed),
		//实现更快的json序列化和更多优化：
		iris.WithOptimizations,
	)

}
