package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

type PageController struct {
	Ctx     iris.Context
	Service services.PageService
}

func (p *PageController) Get() mvc.Result {
	return mvc.View{
		Name: "page/index.html",
		Data: iris.Map{
			"title": "文章列表",
			"data":  "data",
		},
	}
}

func (p *PageController) Post() {

	data := p.Service.GetAll()
	_, _ = p.Ctx.JSON(
		iris.Map{
			"recordsFiltered": 0,
			"recordsTotal":    0,
			"data":            data,
			"start":           0,
		})
}

func (p *PageController) GetBy(page string) mvc.Result {
	data := p.Service.Get(page)
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
		Name: "page.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
