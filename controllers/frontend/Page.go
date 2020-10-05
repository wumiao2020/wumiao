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
	newProduct := product.GetList(6, 0)
	topProduct := product.GetTopList(6, 0)
	news := news.GetList(6, 0)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"title":      "",
			"newProduct": newProduct,
			"topProduct": topProduct,
			"news":       news,
		},
	}
}

func (p *PageController) GetBy(page string) mvc.Result {

	if page == "about.html" {
		return mvc.View{
			Name: "about.html",
		}
	}
	if page == "contact.html" {
		return mvc.View{
			Name: "contact.html",
		}
	}
	if page == "service.html" {
		return mvc.View{
			Name: "service.html",
		}
	}
	if page == "blog.html" {
		return mvc.View{
			Name: "blog.html",
		}
	}
	if page == "shop.html" {
		return mvc.View{
			Name: "shop.html",
		}
	}
	if page == "index.html" {
		return mvc.View{
			Name: "index.html",
		}
	}
	if page == "product-single.html" {
		return mvc.View{
			Name: "product-single.html",
		}
	}
	if page == "blog-single.html" {
		return mvc.View{
			Name: "blog-single.html",
		}
	}

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
