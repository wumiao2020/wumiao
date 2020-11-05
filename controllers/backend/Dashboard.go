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
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *DashboardController) GetBy(page string) mvc.Result {
	data := p.Service.GetByUuid(page)
	if data == nil {
		return mvc.View{
			Code:   iris.StatusNotFound,
			Name:   "errors/404.html",
			Layout: iris.NoLayout,
			Data: iris.Map{
				"title": "你很神，找到了不存在的页面",
			},
		}
	}
	return mvc.View{
		Name: "page/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
