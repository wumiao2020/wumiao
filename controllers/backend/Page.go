package backend

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/models"
	"wumiao/services"
)

type PageController struct {
	Ctx     iris.Context
	Service services.PageService
}

func (p *PageController) Get() mvc.Result {
	return mvc.View{
		Name: "page/index.html",
		Data: iris.Map{
			"title": p.Ctx.Tr("List page"),
			"data":  "data",
		},
	}
}

func (p *PageController) GetCreate() mvc.Result {
	data := new(models.Page)
	return mvc.View{
		Name:   "page/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *PageController) PostSave() {
	data := new(models.Page)
	err := p.Ctx.ReadJSON(&data)
	if len(data.Uuid) == 0 {
		data.Uuid = uuid.NewString()
	}
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

func (p *PageController) PostBy(id int64) {
	dataInfo := p.Service.GetById(id)
	status := 1
	if dataInfo.Status == 1 {
		status = 0
	}
	dataInfo.Status = status
	dataInfo.Id = id
	err := p.Service.Update(dataInfo)
	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": p.Ctx.Tr("Save success !!!"), "id": dataInfo.Id})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
	}
	return
}

func (p *PageController) Post() {

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

func (p *PageController) GetBy(id int64) mvc.Result {
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
		Name:   "page/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"data": data,
		},
	}
}

func (p *PageController) DeleteBy(id int64) {
	err := p.Service.DeleteByID(id)

	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": p.Ctx.Tr("Save success !!!")})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
	}
}
