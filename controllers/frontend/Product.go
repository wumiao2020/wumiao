package frontend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/services"
)

var category = services.NewCategoryService()

type ProductController struct {
	Ctx     iris.Context
	Service services.ProductService
}

func (p *ProductController) Get() mvc.Result {
	categoryData := category.GetByIdentifier("china")
	data := p.Service.GetAll()
	return mvc.View{
		Name: "show/list.html",
		Data: iris.Map{
			"title":    "页面列表",
			"category": categoryData,
			"data":     data,
		},
	}
}

func (p *ProductController) GetBy(product string) mvc.Result {

	if len(product) != 36 {
		categoryData := category.GetByIdentifier(product)
		if categoryData != nil {
			data := p.Service.GetCategoryList(categoryData.Id, 20, 0)
			return mvc.View{
				Name: "show/list.html",
				Data: iris.Map{
					"title":    "页面列表",
					"category": categoryData,
					"data":     data,
				},
			}
		}
	}

	data := p.Service.GetByUuid(product)
	if data == nil {
		return mvc.View{
			Code: iris.StatusNotFound,
			Name: "errors/404.html",
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
