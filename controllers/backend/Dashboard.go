package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

type DashboardController struct {
	Ctx     iris.Context
	Service services.PageService
}

func (p *DashboardController) Get() mvc.Result {
	return mvc.View{
		Name: "dashboard/index.html",
		Data: iris.Map{
			"title": "仪表盘",
			"data":  "data",
		},
	}
}
