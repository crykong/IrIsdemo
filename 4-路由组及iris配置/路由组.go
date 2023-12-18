package main

import (
	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.New()

	//用户模块
	//xxx/users/register 注册
	//xxx/users/login 登录
	//xxx/users/info 获取用户信息

	/*
		路由请求
	*/
	userParty := app.Party("/users", func(context iris.Context) {
		//处理下一个请求
		context.Next()
	})
	/**
	*
	*路由组下面的下一级请求
	*../users/register
	 */
	userParty.Get("register", func(context iris.Context) {
		app.Logger().Info("用户注册")
		context.HTML("<h1>用户组注册功能</h1>")
	})

	userParty.Get("login", func(context iris.Context) {
		app.Logger().Info("用户登录")
		context.HTML("<h1>用户登录功能</h1>")
	})

	userRouter := app.Party("/admin", userMiddleware)

	//done
	userRouter.Done(func(context iris.Context) {
		context.Application().Logger().Info("respose sent to")
	})

	userRouter.Get("info", func(context iris.Context) {
		context.HTML("<h1>用户信息</h1>")
	})
	userRouter.Get("query", func(context iris.Context) {
		context.HTML("<h1>查询信息</h1>")
	})
	userRouter.Done(func(context iris.Context) {
		context.Application().Logger().Info("respose sent to")
	})

	app.Run(iris.Addr(":8090"))
}

func userMiddleware(conent iris.Context) {
	conent.Next()
}
