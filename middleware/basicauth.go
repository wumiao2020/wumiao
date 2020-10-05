package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func Authentication(ctx iris.Context, session *sessions.Session) mvc.Result {

	path := ctx.Path()
	userID := session.GetInt64Default("admin_session_id", 0)
	if userID == 0 {
		if ctx.IsAjax() {
			_, _ = ctx.JSON(iris.Map{"status": false, "code": 401, "message": "请先登录"})

		}

		return mvc.Response{
			// 重定向.
			Path: "/account/login",
		}
	}

	ctx.ViewData("navActive", path)

	ctx.Next() // execute the next handler, in this case the main one.
	return nil
}
