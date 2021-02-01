package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type PageService interface {
	GetAll() []models.Page
	GetList(limit int, start int) []models.Page
	Get(string string) *models.Page
	GetByUuid(string string) *models.Page
	GetById(id int64) *models.Page
	DeleteByID(id int64) error
	Update(data *models.Page) error
	Create(data *models.Page) error
}

type pageService struct {
	engine *xorm.EngineGroup
}

func NewPageService() PageService {
	db := datasource.GetMysqlGroup()
	return &pageService{
		engine: db,
	}
}

func (p pageService) GetList(limit int, start int) []models.Page {
	datalist := make([]models.Page, 0)
	err := p.engine.Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p pageService) GetAll() []models.Page {
	datalist := make([]models.Page, 0)
	err := p.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (p pageService) GetByUuid(string string) *models.Page {
	data := &models.Page{Uuid: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p pageService) GetById(id int64) *models.Page {
	data := &models.Page{Id: id}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}
func (p pageService) DeleteByID(id int64) error {
	data := models.Page{Id: id}
	_, err := p.engine.Id(data.Id).Delete(data)
	return err
}
func (p pageService) Get(string string) *models.Page {
	data := &models.Page{Identifier: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p pageService) Update(data *models.Page) error {
	column := []string{"status"}
	_, err := p.engine.Where("id=?", data.Id).MustCols(column...).Update(data)
	return err
}
func (p pageService) Create(data *models.Page) error {
	_, err := p.engine.Insert(data)
	return err
}
