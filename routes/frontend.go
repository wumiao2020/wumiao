package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"time"
	"wumiao/config"
	"wumiao/controllers/frontend"
	"wumiao/models"
	"wumiao/services"
)

type Menu struct {
	Title      string
	Identifier string
}

func FrontendStart() {
	app := iris.New()

	app.Logger().SetLevel("debug")
	app.HandleDir("/", "./public")
	// 设置关注的视图目录，和文件后缀
	tmpl := iris.HTML("./views/frontend", ".html")
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

	app.Use(func(ctx iris.Context) {
		ctx.ViewData("nav", Nav())
		ctx.Next()
	})

	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	app.RegisterView(tmpl)

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

	//app.Get("/login",controllers.Login)
	//app.Get("/register",controllers.Register)
	err := app.Run(
		iris.Addr(":"+config.GetEnv("FRONTEND_HOST_PORT", "8092")),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	println(err)
}

func notFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}

func Nav() []models.Menu {
	pageService := services.NewMenuService()
	return pageService.GetAll()
}
