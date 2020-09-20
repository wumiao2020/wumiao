package backend

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

type AccountController struct {
	Ctx     iris.Context
	Service services.PageService
}

func (p *AccountController) GetLogin() mvc.Result {
	return mvc.View{
		Layout: "layouts/account.html",
		Name:   "account/login.html",
		Data: iris.Map{
			"title": "登录",
			"data":  "login",
		},
	}
}

func (p *AccountController) PostLogin() {
	f := p.Ctx.FormValues()
	for a, b := range f {
		fmt.Println(a)
		fmt.Println(b)
	}
}

func (p *AccountController) GetRegister() mvc.Result {
	return mvc.View{
		Layout: "layouts/account.html",
		Name:   "account/register.html",
		Data: iris.Map{
			"title": "注册",
			"data":  "register",
		},
	}
}

func (p *AccountController) PostRegister() {
	f := p.Ctx.FormValues()
	for a, b := range f {
		fmt.Println(a)
		fmt.Println(b)
	}
}
