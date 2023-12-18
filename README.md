# IrIsdemo
掌握Go语言Iris框架的使用规范，并能够掌握独自开发web应用的技能。
iris 是一款GO语言种用来开发web应用的框架，该框架支持编写一次在任何地方一最小的机器功率
运行，如Android、Ios、Linx 和window，该框架只需要一个可执行的服务器就可以在平台上运行。
该框架具备一些特点和框架特性 列举如下：
1、聚集搞性能
2、见状的静态路由支持和通配符子域名
3、试图系统支持超过5 以上模板
4、支持定制时间的高可拓展性Websocket API
5、带有GC,内存 redis 提供支持的会话
6、方便的中间和插件
7、完整Rest API
8、能定制HTTP 错误
9、源码改后自动加载。

https://pkg.go.dev/github.com/kataras/iris/v12#section-readme

Installation
Create a new project

$ mkdir IrIsdemo
$ cd IrIsdemo
$ go mod init IrIsdemo
$ go get github.com/kataras/iris/v12@latest # or @v12.2.8