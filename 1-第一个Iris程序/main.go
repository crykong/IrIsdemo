package main

import "github.com/kataras/iris/v12"

func main() {
	//1、创建APP 结构对象
	app := iris.New()
	//2、监听端口
	app.Run(iris.Addr(":9090"), iris.WithoutServerError(iris.ErrServerClosed))
	//app.Run(iris.Addr(":8080"))
}
