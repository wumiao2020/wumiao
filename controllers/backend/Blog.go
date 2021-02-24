package backend

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"html/template"
	"wumiao/models"
	"wumiao/services"
)

type BlogController struct {
	Ctx     iris.Context
	Service services.BlogService
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
		Name: "blog/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *BlogController) PostCreate() {
	postUuid := p.Ctx.PostValueDefault("uuid", "")
	title := p.Ctx.PostValue("title")
	status := p.Ctx.PostValueIntDefault("status", 0)
	content := p.Ctx.FormValue("content")
	identifier := p.Ctx.PostValueDefault("identifier", "")
	metaTitle := p.Ctx.PostValue("meta_title")
	metaKeywords := p.Ctx.PostValue("meta_keywords")
	metaDescription := p.Ctx.PostValue("meta_description")
	thumb := p.Ctx.PostValue("thumb")
	data := models.Blog{Identifier: identifier, Thumb: thumb, MetaTitle: metaTitle, MetaKeywords: metaKeywords, MetaDescription: metaDescription, Title: title, Status: status, Content: template.HTML(content)}
	if postUuid == "" {
		data.Uuid = uuid.New().String()
		if data.Identifier == "" {
			data.Identifier = data.Uuid
		}
		err := p.Service.Create(&data)
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "保存成功！！！", "uuid": data.Uuid})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	} else {
		data.Uuid = postUuid
		if data.Identifier == "" {
			data.Identifier = postUuid
		}
		err := p.Service.Update(&data, []string{"title", "status", "thumb", "content", "identifier", "meta_title", "meta_keywords", "meta_description"})
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "修改成功！！！", "uuid": data.Uuid})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	}

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

func (p *BlogController) GetBy(blog string) mvc.Result {
	data := p.Service.GetByUuid(blog)
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
		Name: "blog/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
