package frontend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

type BlogController struct {
	Ctx     iris.Context
	Service services.BlogService
}

func (p *BlogController) Get() mvc.Result {
	data := p.Service.GetAll()
	return mvc.View{
		Name: "news/list.html",
		Data: iris.Map{
			"title": p.Ctx.Tr("List page"),
			"data":  data,
		},
	}
}

func (p *BlogController) GetBy(news string) mvc.Result {
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
