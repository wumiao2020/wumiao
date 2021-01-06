package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/models"
	"wumiao/services"
)

type RolesController struct {
	Ctx     iris.Context
	Service services.RoleService
}

func (p *RolesController) Get() mvc.Result {
	return mvc.View{
		Name: "page/index.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *RolesController) GetCreate() mvc.Result {
	data := new(models.AdminRoles)
	return mvc.View{
		Name: "page/form.html",
		Data: iris.Map{
			"title": data.Name,
			"data":  data,
		},
	}
}

func (p *RolesController) PostCreate() {

}

func (p *RolesController) Post() mvc.Result {
	//dataAll := p.Service.GetAll()
	data := p.Service.GetAll()
	return mvc.View{
		Layout: iris.NoLayout,
		Name:   "page/list.html",
		Data: iris.Map{
			"data": data,
		},
	}
}

func (p *RolesController) GetBy(id int64) mvc.Result {
	data := p.Service.Get(id)
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
		Name: "page/form.html",
		Data: iris.Map{
			"title": data.Name,
			"data":  data,
		},
	}
}
