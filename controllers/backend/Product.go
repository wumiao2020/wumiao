package backend

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"html/template"
	"wumiao/models"
	"wumiao/services"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.ProductService
}

func (p *ProductController) Get() mvc.Result {
	return mvc.View{
		Name: "product/index.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *ProductController) GetCreate() mvc.Result {
	data := new(models.Product)
	return mvc.View{
		Name: "product/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *ProductController) PostCreate() {
	postUuid := p.Ctx.PostValueDefault("uuid", "")
	title := p.Ctx.PostValue("title")
	price, _ := p.Ctx.PostValueFloat64("price")
	tagPrice, _ := p.Ctx.PostValueFloat64("price")
	position := p.Ctx.PostValueIntDefault("position", 0)
	isActive := p.Ctx.PostValueIntDefault("is_active", 0)
	content := p.Ctx.FormValue("content")
	contentHeading := p.Ctx.FormValue("content_heading")
	identifier := p.Ctx.PostValueDefault("identifier", "")
	metaTitle := p.Ctx.PostValue("meta_title")
	metaKeywords := p.Ctx.PostValue("meta_keywords")
	metaDescription := p.Ctx.PostValue("meta_description")
	thumb := p.Ctx.PostValue("thumb")
	data := models.Product{Identifier: identifier, Price: price, TagPrice: tagPrice, Thumb: thumb, Position: position, MetaTitle: metaTitle, MetaKeywords: metaKeywords, MetaDescription: metaDescription, Title: title, IsActive: isActive, ContentHeading: contentHeading, Content: template.HTML(content)}
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
		err := p.Service.Update(&data, []string{"title", "is_active", "tag_price", "price", "thumb", "position", "content", "content_heading", "identifier", "meta_title", "meta_keywords", "meta_description"})
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "修改成功！！！", "uuid": data.Uuid})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	}

}

func (p *ProductController) Post() {

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

func (p *ProductController) GetBy(product string) mvc.Result {
	data := p.Service.GetByUuid(product)
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
		Name: "product/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
