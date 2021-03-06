package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strings"
)

func Authentication(ctx iris.Context, session *sessions.Session) mvc.Result {

	str := ctx.Application().GetRoutesReadOnly()

	for _, only := range str {
		p := fmt.Sprintf("%v%v", strings.ToLower(only.Method()), strings.ToLower(only.ResolvePath()))
		fmt.Println(p)
		//fmt.Println(strings.Replace(p, "/", ".", -1))
	}

	path := ctx.GetCurrentRoute().ResolvePath()
	//fmt.Println(path)
	errors := session.GetFlash("errors")
	ctx.ViewData("errors", errors)
	userID := session.GetInt64Default("admin_session_id", 0)
	if userID == 0 && path != "/account/login" {
		if ctx.IsAjax() {
			_, _ = ctx.JSON(iris.Map{"status": false, "code": 401, "message": "请先登录"})
		}
		requestUrl := ctx.GetCurrentRoute().Path()
		session.Set("request_url", requestUrl)
		return mvc.Response{
			// 重定向.
			Path: "/account/login",
		}
	}

	ctx.Next() // execute the next handler, in this case the main one.
	return nil
}
