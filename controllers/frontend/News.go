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
			Code: iris.StatusNotFound,
			Name: "errors/404.html",
			Data: iris.Map{
				"title": p.Ctx.Tr("You are very god, found a page that does not exist"),
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
