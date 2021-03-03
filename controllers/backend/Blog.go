package backend

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"wumiao/models"
	"wumiao/services"
)

type BlogController struct {
	Ctx     iris.Context
	Service services.BlogService
	Session *sessions.Session
}

func (p *BlogController) getCurrentUser() *models.Admins {
	user := p.Session.Get(adminSession)
	admins := user.(*models.Admins)
	return admins
}

func (p *BlogController) Get() mvc.Result {
	return mvc.View{
		Name: "blog/index.html",
		Data: iris.Map{
			"title": p.Ctx.Tr("List page"),
			"data":  "data",
		},
	}
}

func (p *BlogController) GetCreate() mvc.Result {
	data := new(models.Blog)
	return mvc.View{
		Name:   "blog/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *BlogController) PostSave() {
	data := new(models.Blog)
	err := p.Ctx.ReadJSON(&data)
	admins := p.getCurrentUser()
	data.Author = admins.Name
	data.AuthorId = admins.Id
	if len(data.Uuid) == 0 {
		data.Uuid = uuid.NewString()
	}
	if len(data.Identifier) == 0 {
		data.Identifier = data.Uuid
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

func (p *BlogController) Post() {

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

func (p *BlogController) GetBy(id int64) mvc.Result {
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
		Name:   "blog/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *BlogController) PostBy(id int64) {
	dataInfo := p.Service.GetById(id)
	status := 1
	if dataInfo.Status == status {
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
