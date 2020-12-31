package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/models"
	"wumiao/services"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.AdminService
}

func (p *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *AdminController) GetCreate() mvc.Result {
	data := new(models.Admins)
	return mvc.View{
		Name:   "admin/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"title": data.Name,
			"data":  data,
		},
	}
}

func (p *AdminController) PostCreate() {
	name := p.Ctx.PostValue("name")
	isActive := p.Ctx.PostValueIntDefault("is_active", 0)
	data := models.Admins{Name: name, Status: isActive}
	err := p.Service.Create(&data)

	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "保存成功！！！", "uuid": data.Id})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
	}

}

func (p *AdminController) Post() {

	limit := p.Ctx.PostValueIntDefault("length", 10)
	start := p.Ctx.PostValueIntDefault("start", 0)

	dataAll := p.Service.GetAll()
	data := p.Service.GetList(limit, start)
	_, _ = p.Ctx.JSON(
		iris.Map{
			"recordsFiltered": len(dataAll),
			"recordsTotal":    len(dataAll),
			"data":            data,
			"start":           0,
		})
}

func (p *AdminController) GetBy(id int64) mvc.Result {
	data := p.Service.GetById(id)
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
		Name: "admin/form.html",
		Data: iris.Map{
			"title": data.Name,
			"data":  data,
		},
	}
}

func (p *AdminController) DeleteBy(id int64) {
	err := p.Service.DeleteByID(id)

	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": p.Ctx.Tr("Save success !!!")})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
	}
}
