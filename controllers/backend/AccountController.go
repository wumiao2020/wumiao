package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"golang.org/x/crypto/bcrypt"
	"wumiao/models"
	"wumiao/services"
)

const adminSessionId = "admin_session_id"

type AccountController struct {
	Ctx     iris.Context
	Service services.AdminService
	// Session，使用依赖注入绑定 main.go.
	Session *sessions.Session
}

func (a *AccountController) GetLogin() mvc.Result {
	if a.isLoggedIn() {

		return mvc.Response{
			// 重定向.
			Path: "/account",
		}

	}
	return mvc.View{
		Layout: "layouts/account.html",
		Name:   "account/login.html",
		Data: iris.Map{
			"title": "登录",
			"data":  "login",
		},
	}
}

func (a *AccountController) PostLogin() {

	username := a.Ctx.PostValue("username")
	password := a.Ctx.PostValue("password")

	list := a.Service.GetList()
	println(list)
	admin := a.Service.GetByEmail(username)
	if admin == nil {
		_, _ = a.Ctx.JSON(iris.Map{"status": false, "message": "用户不存在，请先注册！！！"})
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)) //验证（对比）
		if err != nil {
			_, _ = a.Ctx.JSON(iris.Map{"status": false, "message": "用户名或密码不正确，请重新输入！！！"})
		} else {
			a.Session.Set(adminSessionId, admin.Id)
			_, _ = a.Ctx.JSON(iris.Map{"status": true, "message": "登录成功！！！"})
		}
	}
}

func (a *AccountController) Get() mvc.Result {

	if !a.isLoggedIn() {
		return mvc.Response{
			// 重定向
			Path: "/account/login",
		}
	}
	u := a.Service.GetById(a.getCurrentUserID())
	return mvc.View{
		Layout: "layouts/layout.html",
		Name:   "account/profile.html",
		Data: iris.Map{
			"title": "个人资料",
			"data":  u,
		},
	}
}

func (a *AccountController) GetRegister() mvc.Result {
	return mvc.View{
		Layout: "layouts/account.html",
		Name:   "account/register.html",
		Data: iris.Map{
			"title": "注册",
			"data":  "register",
		},
	}
}

func (a *AccountController) PostRegister() mvc.Result {

	username := a.Ctx.PostValue("username")
	password := a.Ctx.PostValue("password")

	admin := a.Service.GetByEmail(username)
	if admin == nil {
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
		data := models.Admins{Name: username, Email: username, Password: string(hash)}
		_ = a.Service.Create(&data)
		a.Session.Set(adminSessionId, data.Id)
		return mvc.Response{
			Path: "/account",
		}
	}
	return mvc.Response{
		// 重定向
		Path: "/account",
	}
}

func (a *AccountController) AnyLogout() {
	if a.isLoggedIn() {
		a.Session.Destroy()
	}

	a.Ctx.Redirect("/account/login")
}

func (a *AccountController) isLoggedIn() bool {
	return a.getCurrentUserID() > 0
}

func (a *AccountController) getCurrentUserID() int64 {
	userID := a.Session.GetInt64Default(adminSessionId, 0)
	return userID
}
