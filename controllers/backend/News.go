package backend

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/models"
	"wumiao/services"
)

type NewsController struct {
	Ctx     iris.Context
	Service services.NewsService
}

func (p *NewsController) Get() mvc.Result {
	return mvc.View{
		Name: "news/index.html",
		Data: iris.Map{
			"title": "页面列表",
			"data":  "data",
		},
	}
}

func (p *NewsController) GetCreate() mvc.Result {
	data := new(models.Page)
	return mvc.View{
		Name: "news/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *NewsController) PostCreate() {
	postUuid := p.Ctx.PostValueDefault("uuid", "")
	title := p.Ctx.PostValue("title")
	isActive := p.Ctx.PostValueIntDefault("is_active", 0)
	//content := p.Ctx.FormValue("content")
	//data := models.News{Title: title,IsActive: isActive, Content: template.HTML(content)}
	data := models.News{Title: title, IsActive: isActive}
	if postUuid == "" {
		data.Uuid = uuid.New().String()
		err := p.Service.Create(&data)
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "保存成功！！！", "uuid": data.Uuid})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	} else {
		data.Uuid = postUuid
		err := p.Service.Update(&data, []string{"title", "parent_id", "is_active", "content"})
		if err == nil {
			_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": "修改成功！！！", "uuid": data.Uuid})
		} else {
			_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
		}
	}

}

func (p *NewsController) Post() {
	data := p.Service.GetAll()
	_, _ = p.Ctx.JSON(
		iris.Map{
			"recordsFiltered": 0,
			"recordsTotal":    0,
			"data":            data,
			"start":           0,
		})
}

func (p *NewsController) GetBy(page string) mvc.Result {
	data := p.Service.GetByUuid(page)
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
		Name: "page/form.html",
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}
