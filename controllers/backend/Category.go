package backend

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"html/template"
	"wumiao/models"
	"wumiao/services"
)

type CategoryController struct {
	Ctx     iris.Context
	Service services.CategoryService
}

func (p *CategoryController) Get() mvc.Result {
	return mvc.View{
		Name: "category/index.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *CategoryController) GetCreate() mvc.Result {
	data := new(models.Category)
	return mvc.View{
		Name: "category/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *CategoryController) PostCreate() {
	postUuid := p.Ctx.PostValueDefault("uuid", "")
	title := p.Ctx.PostValue("title")
	isActive := p.Ctx.PostValueIntDefault("is_active", 0)
	content := p.Ctx.FormValue("content")
	identifier := p.Ctx.PostValueDefault("identifier", "")
	metaTitle := p.Ctx.PostValue("meta_title")
	metaKeywords := p.Ctx.PostValue("meta_keywords")
	metaDescription := p.Ctx.PostValue("meta_description")
	thumb := p.Ctx.PostValue("thumb")
	data := models.Category{Identifier: identifier, Thumb: thumb, MetaTitle: metaTitle, MetaKeywords: metaKeywords, MetaDescription: metaDescription, Title: title, IsActive: isActive, Content: template.HTML(content)}
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
		err := p.Service.Update(&data, []string{"title", "is_active", "thumb", "content", "identifier", "meta_title", "meta_keywords", "meta_description"})
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "修改成功！！！", "uuid": data.Uuid})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	}

}

func (p *CategoryController) Post() {

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

func (p *CategoryController) GetBy(category string) mvc.Result {
	data := p.Service.GetByUuid(category)
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
		Name: "category/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
