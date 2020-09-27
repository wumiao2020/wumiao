package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
	"wumiao/config"
	"wumiao/controllers/backend"
	"wumiao/middleware"
	"wumiao/services"
)

var sessManager = sessions.New(sessions.Config{
	Cookie:  "sessioncookiename",
	Expires: 1 * time.Hour,
})

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

	hero.Register(sessManager.Start)

	account := mvc.New(app.Party("/account"))
	adminService := services.NewAdminService()
	account.Register(
		adminService,
	)
	account.Handle(new(backend.AccountController))

	app.Use(hero.Handler(middleware.Authentication))

	app.Use(func(ctx iris.Context) {
		ctx.ViewData("nav", Nav())
		ctx.Next()
	})

	//页面配置
	cms := mvc.New(app.Party("/cms"))
	pageService := services.NewPageService()
	cms.Register(pageService)
	cms.Handle(new(backend.PageController))

	err := app.Run(
		iris.Addr(":"+config.GetEnv("BACKEND_HOST_PORT", "8091")),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	println(err)
}
