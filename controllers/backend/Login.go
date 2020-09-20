package backend

import "github.com/kataras/iris/v12"

func Login(ctx iris.Context) {
	ctx.ViewLayout(iris.NoLayout)
	ctx.ViewData("title", "登录")
	err := ctx.View("login.html")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.Writef(err.Error())
	}
}

func PostLogin(ctx iris.Context) {

}
