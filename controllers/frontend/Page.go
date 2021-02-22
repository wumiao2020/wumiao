package frontend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

var product = services.NewProductService()
var news = services.NewNewsService()

type PageController struct {
	Ctx     iris.Context
	Service services.PageService
}

func (p *PageController) Get() mvc.Result {
	data := p.Service.Get("index.html")

	if data == nil {
		return mvc.View{
			Code:   iris.StatusNotFound,
			Name:   "errors/404.html",
			Layout: iris.NoLayout,
			Data: iris.Map{
				"title": p.Ctx.Tr("You are very god, found a page that does not exist"),
			},
		}
	}

	newProduct := product.GetList(6, 0)
	topProduct := product.GetTopList(6, 0)
	news := news.GetList(6, 0)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"title":      data.Title,
			"newProduct": newProduct,
			"topProduct": topProduct,
			"news":       news,
			"data":       data,
		},
	}
}

func (p *PageController) GetBy(page string) mvc.Result {

	if page == "contact.html" {
		return mvc.View{
			Name: "contact.html",
		}
	}

	data := p.Service.Get(page)

	if data == nil {
		return mvc.View{
			Code:   iris.StatusNotFound,
			Name:   "errors/404.html",
			Layout: iris.NoLayout,
			Data: iris.Map{
				"title": p.Ctx.Tr("You are very god, found a page that does not exist"),
			},
		}
	}

	if page == "index.html" {
		newProduct := product.GetList(6, 0)
		topProduct := product.GetTopList(6, 0)
		news := news.GetList(6, 0)
		return mvc.View{
			Name: "index.html",
			Data: iris.Map{
				"title":      data.Title,
				"newProduct": newProduct,
				"topProduct": topProduct,
				"news":       news,
				"data":       data,
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
