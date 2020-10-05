package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/models"
	"wumiao/services"
)

type MenuController struct {
	Ctx     iris.Context
	Service services.MenuService
}

func (p *MenuController) Get() mvc.Result {
	return mvc.View{
		Name: "menu/index.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *MenuController) GetCreate() mvc.Result {
	data := new(models.Menu)
	return mvc.View{
		Name: "menu/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *MenuController) PostCreate() {
	id := p.Ctx.PostValueInt64Default("id", 0)
	title := p.Ctx.PostValue("title")
	uri := p.Ctx.PostValue("uri")
	isActive := p.Ctx.PostValueIntDefault("is_active", 0)
	position := p.Ctx.PostValueIntDefault("position", 0)
	data := models.Menu{Title: title, IsActive: isActive, Position: position, Uri: uri}
	if id == 0 {
		err := p.Service.Create(&data)
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "保存成功！！！", "uuid": data.Id})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	} else {
		data.Id = id
		err := p.Service.Update(&data, []string{"title", "position", "uri", "is_active", "thumb", "identifier"})
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "修改成功！！！", "uuid": data.Id})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	}

}

func (p *MenuController) Post() {

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

func (p *MenuController) GetBy(id int64) mvc.Result {
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
		Name: "menu/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
