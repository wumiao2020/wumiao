package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
	"wumiao/config"
	"wumiao/controllers/backend"
	"wumiao/services"
)

func BackendStart() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.HandleDir("/", "./public/material")
	// 设置关注的视图目录，和文件后缀
	tmpl := iris.HTML("./views/backend", ".html")
	tmpl.Layout("layouts/layout.html")
	// 是否每次请求都重新加载文件，这个在开发期间设置为true，在发布时设置为false
	// 可以方便每次修改视图文件而无需停止服务
	tmpl.Reload(true)

	tmpl.AddFunc("greet", func(x int, y int) bool {
		return (x+1)%y == 0
	})

	tmpl.AddFunc("nowYear", func() int {
		return time.Now().UTC().Year()
	})

	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	app.RegisterView(tmpl)
	// "/user" 基于mvc的应用程序.
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookiename",
		Expires: 24 * time.Hour,
	})

	account := mvc.New(app.Party("/account"))
	adminService := services.NewAdminService()
	account.Register(
		adminService,
		sessManager.Start,
	)
	account.Handle(new(backend.AccountController))

	app.Use(before)

	app.Use(func(ctx iris.Context) {
		ctx.ViewData("nav", Nav())
		ctx.Next()
	})

	page := mvc.New(app.Party("/"))
	pageService := services.NewPageService()
	page.Register(pageService)
	page.Handle(new(backend.PageController))
	err := app.Run(
		iris.Addr(":"+config.GetEnv("BACKEND_HOST_PORT", "8091")),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	println(err)
}

func before(ctx iris.Context) {

	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // execute the next handler, in this case the main one.
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}
