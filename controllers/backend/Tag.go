package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"html/template"
	"wumiao/models"
	"wumiao/services"
)

type TagController struct {
	Ctx     iris.Context
	Service services.TagService
}

func (p *TagController) Get() mvc.Result {
	return mvc.View{
		Name: "tag/index.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *TagController) GetCreate() mvc.Result {
	data := new(models.Page)
	return mvc.View{
		Name: "tag/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *TagController) PostCreate() {
	postUuid := p.Ctx.PostValueDefault("uuid", "")
	title := p.Ctx.PostValue("title")
	content := p.Ctx.FormValue("content")
	data := models.Tag{Title: title, Content: template.HTML(content)}
	if postUuid == "" {
		err := p.Service.Create(&data)
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "保存成功！！！"})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	} else {
		err := p.Service.Update(&data, []string{"title", "parent_id", "is_active", "content"})
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "修改成功！！！"})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	}

}

func (p *TagController) Post() {
	data := p.Service.GetAll()
	_, _ = p.Ctx.JSON(
		iris.Map{
			"recordsFiltered": 0,
			"recordsTotal":    0,
			"data":            data,
			"start":           0,
		})
}

func (p *TagController) GetBy(id int) mvc.Result {
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
		Name: "tag/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
