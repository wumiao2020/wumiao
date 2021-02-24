package backend

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"wumiao/models"
	"wumiao/services"
)

type BlockController struct {
	Ctx     iris.Context
	Service services.BlockService
}

func (p *BlockController) Get() mvc.Result {
	return mvc.View{
		Name: "block/index.html",
		Data: iris.Map{
			"title": p.Ctx.Tr("List block"),
			"data":  "data",
		},
	}
}

func (p *BlockController) GetCreate() mvc.Result {
	data := new(models.Block)
	return mvc.View{
		Name:   "block/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"title": data.Title,
			"data":  data,
		},
	}
}

func (p *BlockController) PostSave() {
	data := new(models.Block)
	err := p.Ctx.ReadJSON(&data)
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

func (p *BlockController) PostBy(id int64) {
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

func (p *BlockController) Post() {

	limit := p.Ctx.PostValueIntDefault("length", 10)
	search := p.Ctx.PostValue("search[value]")
	start := p.Ctx.PostValueIntDefault("start", 0)

	dataAll := p.Service.GetAll(search)
	data := p.Service.GetList(limit, start, search)
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

func (p *BlockController) GetBy(id int64) mvc.Result {
	data := p.Service.GetById(id)
	if data == nil {
		return mvc.View{
			Code:   iris.StatusNotFound,
			Name:   "errors/404.html",
			Layout: iris.NoLayout,
			Data: iris.Map{
				"title": p.Ctx.Tr("You are very god, found a block that does not exist"),
			},
		}
	}
	return mvc.View{
		Name:   "block/form.html",
		Layout: iris.NoLayout,
		Data: iris.Map{
			"data": data,
		},
	}
}

func (p *BlockController) DeleteBy(id int64) {
	err := p.Service.DeleteByID(id)

	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{"status": true, "message": p.Ctx.Tr("Save success !!!")})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{"status": false, "message": err})
	}
}
