package backend

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"wumiao/models"
	"wumiao/services"
)

type ApiController struct {
	Ctx     iris.Context
	Service services.ApiService
}

func (p *ApiController) Get() {
	dataAll := p.Service.GetAll()
	_, _ = p.Ctx.JSON(
		iris.Map{
			"status": true,
			"code":   200,
			"data":   dataAll,
		})
}

func (p *ApiController) Post() {
	data := new(models.Api)
	e := p.Ctx.ReadJSON(&data)
	fmt.Println(e)
	data.Number = 1
	err := p.Service.Create(data)

	if err == nil {
		_, _ = p.Ctx.JSON(iris.Map{
			"code": 200, "status": true, "message": "Save success！！！", "id": data.Id})
	} else {
		_, _ = p.Ctx.JSON(iris.Map{
			"code": 500, "status": false, "message": err})
	}

}

func (p *ApiController) DeleteBy(id int64) {
	err := p.Service.DeleteByID(id)

	if err == nil {
		_, err = p.Ctx.JSON(iris.Map{
			"code": 200, "status": true, "message": "Delete success !!!"})
		fmt.Println(err)
	} else {
		_, err = p.Ctx.JSON(iris.Map{
			"code": 500, "status": false, "message": err})
		fmt.Println(err)
	}
}
