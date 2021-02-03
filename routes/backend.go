package routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strings"
	"time"
	"wumiao/config"
	"wumiao/controllers/backend"
	"wumiao/middleware"
	"wumiao/models"
	"wumiao/services"
)

var sessManager = sessions.New(sessions.Config{
	Cookie:  "sessioncookiename",
	Expires: 1 * time.Hour,
})

func BackendStart() {

	app := iris.New()

	err := app.I18n.Load("./locales/*/*.ini", "en-US", "zh-TW", "zh-CN")

	app.I18n.SetDefault("zh-CN")
	//
	//app.I18n.URLParameter = "l"
	//app.I18n.ExtractFunc = func(ctx iris.Context) string {
	//
	//	language := ctx.URLParam("l")
	//	if len(language) > 0 {
	//		ctx.RemoveCookie("language")
	//		ctx.SetCookieKV("language",language)
	//	}else {
	//		language = ctx.GetCookie("language")
	//	}
	//
	//	println(language)
	//
	//	switch language {
	//	case "tw":
	//		language = "zh-TW"
	//	case "en":
	//		language = "en-US"
	//	case "cn":
	//		language = "zh-CN"
	//	default :
	//		language = "zh-CN"
	//	}
	//
	//
	//	return language // if empty then it will continue with the rest.
	//}
	fmt.Println(err)
	//app.Logger().SetLevel("debug")
	app.HandleDir("/argon", "./public/argon")
	app.HandleDir("/assets", "./public/argon/assets")
	app.HandleDir("/upload", "./public/upload")
	// 设置关注的视图目录，和文件后缀
	tmpl := iris.HTML("./views/backend", ".html")
	tmpl.Layout("layouts/layout.html")
	// 是否每次请求都重新加载文件，这个在开发期间设置为true，在发布时设置为false
	// 可以方便每次修改视图文件而无需停止服务
	tmpl.Reload(true)

	app.OnErrorCode(iris.StatusNotFound, backendNotFound)
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

	account := mvc.New(app.Party("/account"))
	accountService := services.NewAdminService()
	account.Register(accountService)
	account.Handle(new(backend.AccountController))

	app.Use(hero.Handler(middleware.Authentication))

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

	//管理员管理
	admin := mvc.New(app.Party("/admin"))
	adminService := services.NewAdminService()
	admin.Register(adminService)
	admin.Handle(new(backend.AdminController))

	//角色管理
	role := mvc.New(app.Party("/role"))
	roleService := services.NewAdminRolesService()
	role.Register(roleService)
	role.Handle(new(backend.AdminRoleController))

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

	//标签管理
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

	err = app.Run(
		iris.Addr(":"+config.GetEnv("BACKEND_HOST_PORT", "8091")),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	fmt.Println(err)
}

func backendNotFound(ctx iris.Context) {
	// 出现 404 的时候，就跳转到 $views_dir/errors/404.html 模板
	ctx.ViewLayout("layouts/account.html")
	ctx.ViewData("data", "")
	_ = ctx.View("errors/404.html")
}

func BackendHtml() {

	app := iris.New()
	//app.Logger().SetLevel("debug")
	app.HandleDir("/", "./public/freedom")

	err := app.Run(
		iris.Addr(":8080"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	fmt.Println(err)
}

func internalServerError(ctx iris.Context) {
	_, _ = ctx.WriteString("Oups something went wrong, try again")
}

func MenuList() []models.AdminPermissions {
	permissionService := services.NewPermissionService()
	return permissionService.GetMenuList()
}

func Breadcrumbs(path string) []models.AdminPermissions {
	permissionService := services.NewPermissionService()
	return permissionService.GetBreadcrumbs(path)
}

func breadcrumbBut(perms string) *models.AdminPermissions {
	permissionService := services.NewPermissionService()
	return permissionService.GetBreadcrumbBut(perms)
}
