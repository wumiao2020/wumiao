package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/models"
	"wumiao/services"
)

type PermissionController struct {
	Ctx     iris.Context
	Service services.PermissionService
}

func (p *PermissionController) Get() mvc.Result {
	return mvc.View{
		Name: "permission/index.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *PermissionController) GetCreate() mvc.Result {
	data := new(models.AdminPermissions)
	return mvc.View{
		Name: "permission/form.html",
		Data: iris.Map{
			"title": data.Name,
			"data":  data,
		},
	}
}

func (p *PermissionController) PostCreate() {

}

func (p *PermissionController) Post() mvc.Result {
	data := p.Service.GetAll()
	return mvc.View{
		Layout: iris.NoLayout,
		Name:   "permission/list.html",
		Data: iris.Map{
			"data": data,
		},
	}
}

func (p *PermissionController) GetBy(uuid int64) mvc.Result {
	data := p.Service.Get(uuid)
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
		Name: "permission/form.html",
		Data: iris.Map{
			"title": data.Name,
			"data":  data,
		},
	}
}
