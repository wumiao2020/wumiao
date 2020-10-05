package frontend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.ProductService
}

func (p *ProductController) Get() mvc.Result {
	data := p.Service.GetAll()
	return mvc.View{
		Name: "show/list.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  data,
		},
	}
}

func (p *ProductController) GetBy(product string) mvc.Result {
	data := p.Service.GetByUuid(product)
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
		Name: "show/single.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
