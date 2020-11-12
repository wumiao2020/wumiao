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
	//app.Logger().SetLevel("debug")
	app.HandleDir("/assets", "./public/admin/assets")
	app.HandleDir("/upload", "./public/upload")
	// 设置关注的视图目录，和文件后缀
	tmpl := iris.HTML("./views/backend", ".html")
	tmpl.Layout("layouts/layout.html")
	// 是否每次请求都重新加载文件，这个在开发期间设置为true，在发布时设置为false
	// 可以方便每次修改视图文件而无需停止服务
	tmpl.Reload(true)

	tmpl.AddFunc("isActionMenu", func(x string, y string) bool {
		return x == y
	})

	tmpl.AddFunc("nowYear", func() int {
		return time.Now().UTC().Year()
	})

	app.OnErrorCode(iris.StatusNotFound, backendNotFound)
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
		path := ctx.Path()
		ctx.ViewData("navActive", path)
		ctx.ViewData("nav", Nav())
		ctx.Next()
	})

	//页面管理
	dashboard := mvc.New(app.Party("/"))
	dashboardService := services.NewPageService()
	dashboard.Register(dashboardService)
	dashboard.Handle(new(backend.DashboardController))

	//权限管理
	permission := mvc.New(app.Party("/permission"))
	permissionService := services.NewPermissionService()
	permission.Register(permissionService)
	permission.Handle(new(backend.PermissionController))

	//角色管理
	role := mvc.New(app.Party("/role"))
	roleService := services.NewRoleService()
	role.Register(roleService)
	role.Handle(new(backend.RolesController))

	//页面管理
	page := mvc.New(app.Party("/page"))
	pageService := services.NewPageService()
	page.Register(pageService)
	page.Handle(new(backend.PageController))
	//文章管理
	news := mvc.New(app.Party("/news"))
	newsService := services.NewNewsService()
	news.Register(newsService)
	news.Handle(new(backend.NewsController))
	//分类管理
	tag := mvc.New(app.Party("/tag"))
	tagService := services.NewTagService()
	tag.Register(tagService)
	tag.Handle(new(backend.TagController))
	//导航管理
	menu := mvc.New(app.Party("/menu"))
	menuService := services.NewMenuService()
	menu.Register(menuService)
	menu.Handle(new(backend.MenuController))
	//产品管理
	product := mvc.New(app.Party("/product"))
	productService := services.NewProductService()
	product.Register(productService)
	product.Handle(new(backend.ProductController))
	//分类管理
	category := mvc.New(app.Party("/category"))
	categoryService := services.NewCategoryService()
	category.Register(categoryService)
	category.Handle(new(backend.CategoryController))
	//上传管理
	upload := mvc.New(app.Party("/upload"))
	upload.Handle(new(backend.UploadController))

	err := app.Run(
		iris.Addr(":"+config.GetEnv("BACKEND_HOST_PORT", "8091")),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	println(err)
}

func backendNotFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	ctx.ViewLayout(iris.NoLayout)
	ctx.View("errors/404.html")
}

func BackendHtml() {

	app := iris.New()
	//app.Logger().SetLevel("debug")
	app.HandleDir("/", "./public/admin")

	err := app.Run(
		iris.Addr(":8080"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	println(err)
}
