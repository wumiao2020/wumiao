package frontend

import "github.com/kataras/iris/v12"

func Index(ctx iris.Context) {
	ctx.ViewData("title", "首页")
	ctx.ViewData("active", "index.html")
	err := ctx.View("index.html")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.Writef(err.Error())
	}
}
