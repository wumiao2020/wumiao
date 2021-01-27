package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type ApiService interface {
	GetAll() []models.Api
	DeleteByID(id int64) error
	Update(data *models.Api, columns []string) error
	Create(data *models.Api) error
}

type apiService struct {
	engine *xorm.EngineGroup
}

func NewApiService() ApiService {
	db := datasource.GetMysqlGroup()
	return &apiService{
		engine: db,
	}
}

func (a apiService) GetAll() []models.Api {
	datalist := make([]models.Api, 0)
	err := a.engine.Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (a apiService) Update(data *models.Api, column []string) error {
	_, err := a.engine.Id(data.Id).MustCols(column...).Update(data)
	return err
}

func (a apiService) DeleteByID(id int64) error {
	data := models.Api{Id: id}
	_, err := a.engine.Id(data.Id).Delete(data)
	return err
}

func (a apiService) Create(data *models.Api) error {
	_, err := a.engine.Insert(data)
	return err
}
