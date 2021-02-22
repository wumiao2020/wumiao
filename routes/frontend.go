package routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
	"strings"
	"time"
	"wumiao/config"
	"wumiao/controllers/backend"
	"wumiao/controllers/frontend"
	"wumiao/models"
	"wumiao/services"
)

func Frontend() {

	app := iris.New()

	err := app.I18n.Load("./locales/frontend/*/*.ini", "en-US", "zh-TW", "zh-CN")

	app.I18n.SetDefault("zh-CN")
	fmt.Println(err)
	//app.Logger().SetLevel("debug")
	app.HandleDir("/assets", "./public/frontend/assets")
	app.HandleDir("/upload", "./public/upload")
	// 设置关注的视图目录，和文件后缀
	tmpl := iris.HTML("./views/frontend", ".html")
	tmpl.Layout("layouts/layout.html")
	// 是否每次请求都重新加载文件，这个在开发期间设置为true，在发布时设置为false
	// 可以方便每次修改视图文件而无需停止服务
	tmpl.Reload(true)

	app.OnErrorCode(iris.StatusNotFound, frontendNotFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	app.RegisterView(tmpl)

	hero.Register(sessManager.Start)

	app.Use(func(ctx iris.Context) {
		perms := strings.Replace(strings.ToLower(ctx.GetCurrentRoute().Name()), "/", ".", -1)
		breadcrumbs := Breadcrumbs(perms)
		Ap := new(models.AdminPermissions)
		if len(breadcrumbs) > 0 {
			breadcrumbBut := breadcrumbBut(perms + ".create")
			breadcrumb := breadcrumbs[len(breadcrumbs)-1]
			ctx.ViewData("breadcrumbBut", breadcrumbBut)
			ctx.ViewData("breadcrumb", breadcrumb)
		} else {
			ctx.ViewData("breadcrumbBut", Ap)
			ctx.ViewData("breadcrumb", Ap)
		}
		ctx.ViewData("breadcrumbs", breadcrumbs)
		ctx.ViewData("perms", perms)
		ctx.ViewData("menuList", MenuList())
		ctx.ViewData("tr", ctx.Tr)
		ctx.Next()
	})

	tmpl.AddFunc("isAction", func(id int64, breadcrumbs []models.AdminPermissions) bool {
		for _, breadcrumb := range breadcrumbs {
			if breadcrumb.Id == id {
				return true
			}
		}
		return false
	})

	tmpl.AddFunc("nowYear", func() int {
		return time.Now().UTC().Year()
	})

	page := mvc.New(app.Party("/"))
	pageService := services.NewPageService()
	page.Register(pageService)
	page.Handle(new(frontend.PageController))

	news := mvc.New(app.Party("/news"))
	newsService := services.NewNewsService()
	news.Register(newsService)
	news.Handle(new(frontend.NewsController))

	shop := mvc.New(app.Party("/shop"))
	shopService := services.NewProductService()
	shop.Register(shopService)
	shop.Handle(new(frontend.ProductController))

	//上传管理
	upload := mvc.New(app.Party("/upload"))
	upload.Handle(new(backend.UploadController))

	err = app.Run(
		iris.Addr(":"+config.GetEnv("FRONTEND_HOST_PORT", "8080")),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	fmt.Println(err)
}

func frontendNotFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	ctx.ViewLayout(iris.NoLayout)
	ctx.ViewData("data", "")
	_ = ctx.View("errors/404.html")
}

func FrontendHtml() {

	app := iris.New()
	//app.Logger().SetLevel("debug")
	app.HandleDir("/", "./public/frontend")

	err := app.Run(
		iris.Addr(":8081"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	fmt.Println(err)
}
