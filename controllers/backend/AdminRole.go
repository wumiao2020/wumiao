package backend

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/models"
	"wumiao/services"
)

type AdminRoleController struct {
	Ctx     iris.Context
	Service services.AdminRolesService
}

func (p *AdminRoleController) Get() mvc.Result {
	return mvc.View{
		Name: "role/index.html",
		Data: iris.Map{
			"title": p.Ctx.Tr("List page"),
			"data":  "data",
		},
	}
}

func (p *AdminRoleController) GetCreate() mvc.Result {
	data := models.Admins{Status: 1}
	return mvc.View{
		Name:   "role/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"title": data.Name,
			"data":  data,
		},
	}
}

func (p *AdminRoleController) PostSave() {
	data := new(models.AdminRoles)
	err := p.Ctx.ReadJSON(&data)
	fmt.Println(err)
	if data.Id == 0 {
		err = p.Service.Create(data)
	} else {
		err = p.Service.Update(data)
	}

	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": p.Ctx.Tr("Save success !!!"), "id": data.Id})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
	}
	return
}

func (p *AdminRoleController) PostBy(id int64) {
	dataInfo := p.Service.GetById(id)
	data := new(models.AdminRoles)
	status := 1
	if dataInfo.Status == 1 {
		status = 0
	}
	data.Status = status
	data.Id = id
	err := p.Service.Update(data)
	fmt.Println(err)
	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": p.Ctx.Tr("Save success !!!"), "id": data.Id})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
	}
	return
}

func (p *AdminRoleController) Post() {

	limit := p.Ctx.PostValueIntDefault("length", 10)
	start := p.Ctx.PostValueIntDefault("start", 0)

	dataAll := p.Service.GetAll()
	data := p.Service.GetList(limit, start)
	_, _ = p.Ctx.JSON(
		iris.Map{
			"status":          false,
			"code":            200,
			"recordsFiltered": len(dataAll),
			"recordsTotal":    len(dataAll),
			"data":            data,
			"start":           0,
		})
}

func (p *AdminRoleController) GetBy(id int64) mvc.Result {
	data := p.Service.GetById(id)
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
		Name:   "role/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"data": data,
		},
	}
}

func (p *AdminRoleController) DeleteBy(id int64) {
	err := p.Service.DeleteByID(id)

	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": p.Ctx.Tr("Save success !!!")})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
	}
}
