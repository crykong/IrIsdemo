package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	/**
	* handle 方式处理请求
	//get:http://localhost:8080/userinfo
	*/

	app.Handle("GET", "/userinfo", func(ctx iris.Context) {
		//获取path
		path := ctx.Path()
		app.Logger().Info(path)
		app.Logger().Error("request Path:", path)
	})

	app.Handle("PSOT", "/postcommit", func(ctx iris.Context) {
		//获取path
		path := ctx.Path()
		app.Logger().Error("post request Path:", path)
		ctx.HTML(path)
	})

	/**
	* 1、 Get 正则表达式 路由
	使用：context.Params().Get("name") 获取正则表达式 变量
	请求1：
	http://localhost:8080/weather/2019-03-11/shanghai
	http://localhost:8080/weather/2019-03-10/beijing
	http://localhost:8080/weather/2019-03-09/chognqing
	*/
	app.Get("/weather/{date}/{city}", func(ctx iris.Context) {
		//获取变量
		path := ctx.Path()
		date := ctx.Params().Get("date")
		city := ctx.Params().Get("city")

		ctx.WriteString(path + " ," + date + ", " + city)
	})

	/**
	* 2、 自定义正则表达式变量路由请求（unit64:uint64） 进行变量限制
	 */
	app.Get("/api/users/{userid:uint64}}", func(ctx iris.Context) {
		userId, err := ctx.Params().GetInt("userid")

		if err != nil {
			ctx.JSON(map[string]interface{}{
				"code":    201,
				"message": "bad request",
			})
			return
		}

		ctx.JSON(map[string]interface{}{
			"code":   200,
			"userid": userId,
		})
	})

	/**
	* 2、 自定义正则表达式变量路由请求bool
	 */
	app.Get("/api/users/{islogin:bool}}", func(ctx iris.Context) {
		islogin, err := ctx.Params().GetBool("islogin")

		if err != nil {
			ctx.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if islogin {
			ctx.WriteString("已经登录")
		} else {
			ctx.WriteString("未登录")
		}
	})

	app.Run(iris.Addr(":8080"))

}
