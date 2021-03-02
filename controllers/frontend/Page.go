package frontend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

var product = services.NewProductService()
var blog = services.NewBlogService()

type PageController struct {
	Ctx     iris.Context
	Service services.PageService
}

func (p *PageController) Get() mvc.Result {

	newProduct := product.GetList(6, 0)
	topProduct := product.GetTopList(6, 0)
	blog := blog.GetList(6, 0)
	return mvc.View{
		Name: "page/index.html",
		Data: iris.Map{
			"title":      p.Ctx.Tr("Home Page"),
			"newProduct": newProduct,
			"topProduct": topProduct,
			"blog":       blog,
		},
	}
}

func (p *PageController) GetBy(page string) mvc.Result {

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
	return mvc.View{
		Name: "page/" + data.PageLayout,
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *PageController) GetContact() mvc.Result {
	return mvc.View{
		Name: "page/contact.html",
		Data: iris.Map{
			"title": p.Ctx.Tr("Contact Us"),
		},
	}
}

func (p *PageController) GetAbout() mvc.Result {
	return mvc.View{
		Name: "page/about.html",
		Data: iris.Map{
			"title": p.Ctx.Tr("About Us"),
		},
	}
}
