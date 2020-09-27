package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func Authentication(ctx iris.Context, session *sessions.Session) mvc.Result {

	userID := session.GetInt64Default("admin_session_id", 0)
	if userID == 0 {
		return mvc.Response{
			// 重定向.
			Path: "/account/login",
		}
	}

	ctx.Values().Set("info", "info")

	ctx.Next() // execute the next handler, in this case the main one.
	return nil
}
