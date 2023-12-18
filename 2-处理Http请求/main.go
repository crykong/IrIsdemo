package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	//ur://http://localhost:8080/getRequest
	//type:get
	app.Get("/getRequest", func(ctx iris.Context) {
		// 处理 GET 请求，请求的URL 为：/getRequest
		path := ctx.Path()
		app.Logger().Info(path)
	})
	//1、处理GET 请求
	app.Get("/userpath", func(ctx iris.Context) {
		//获取path
		path := ctx.Path()
		app.Logger().Info(path)
		ctx.WriteString("请求路径" + path)
	})

	//http://localhost:8080/userinfo?username=zhangsan&pwd=212121
	//2、处理GeT 请求，并且接受参数
	app.Get("/userinfo", func(ctx iris.Context) {
		//获取path
		path := ctx.Path()
		app.Logger().Info(path)

		userName := ctx.URLParam("username")
		app.Logger().Info(userName)
		pwd := ctx.URLParam("pwd")
		app.Logger().Info(pwd)

		//返回HTML 数据
		ctx.HTML("<H1> 用户名称：" + userName + "密码:" + pwd + "</H1>")

	})

	//3、处理POST 请求
	app.Post("/userlogin", func(ctx iris.Context) {
		//获取path
		path := ctx.Path()
		app.Logger().Info(path)
		//用conent.PostValue 方法获取POST 请求提交的From表单数据
		userName := ctx.PostValue("username")
		age := ctx.PostValue("age")
		app.Logger().Info(userName, "  ", age)

		ctx.HTML(userName)
	})

	//4、处理POST Json 数据请求
	app.Post("/postjson", func(ctx iris.Context) {
		//获取path
		path := ctx.Path()
		app.Logger().Info(path)

		var person Person

		if err := ctx.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		ctx.Writef("rececied: %#+v\n", person)

	})
	//3、处理POST XML 数据请求
	/*
		请求配置：Conent-type到application/xml //可选但是可以配置
		*请求内容
		*
		<student>
		<stu_name>zhangsan</stu_name>
		<stu_age>28</stu_age>
		</student>
	*/
	app.Post("/postxml", func(ctx iris.Context) {
		//获取path
		path := ctx.Path()
		app.Logger().Info("请求ULR", path)

		var student Student

		if err := ctx.ReadXML(&student); err != nil {
			panic(err.Error())
		}
		ctx.Writef("rececied: %#+v\n", student)

	})

	//请求返回数据
	//PUT
	app.Put("getHell", func(context iris.Context) {
		context.WriteString("Put Requset")

	})

	//Delete
	app.Delete("deletHell", func(context iris.Context) {
		context.WriteString("Delete Requset")

	})

	app.Get("getjson", func(context iris.Context) {
		context.JSON(iris.Map{"msg": "hello word", "code": 200})
		context.JSON(Person{Name: "张三", Age: 21})
		context.XML(Student{Name: "张三", Age: 21})
		context.Text("Hello world")
	})

	// 运行应用
	app.Run(iris.Addr(":8080"))
}

// `string` on pkg.go.de
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 自定义结构体
type Student struct {
	//XMLANMW xml.name "
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}
