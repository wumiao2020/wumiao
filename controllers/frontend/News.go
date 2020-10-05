package frontend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

type NewsController struct {
	Ctx     iris.Context
	Service services.NewsService
}

func (p *NewsController) Get() mvc.Result {
	data := p.Service.GetAll()
	return mvc.View{
		Name: "news/list.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  data,
		},
	}
}

func (p *NewsController) GetBy(news string) mvc.Result {
	data := p.Service.GetByUuid(news)
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
		Name: "news/single.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
