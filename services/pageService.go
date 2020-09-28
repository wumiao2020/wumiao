package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type PageService interface {
	GetAll() []models.Page
	GetList(parentId int) []models.Page
	Get(string string) *models.Page
	GetByUuid(string string) *models.Page
	DeleteByID(id int64) error
	Update(data *models.Page, columns []string) error
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

func (p pageService) GetList(parentId int) []models.Page {
	datalist := make([]models.Page, 0)
	err := p.engine.Where("parent_id=?", parentId).Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p pageService) GetAll() []models.Page {
	datalist := make([]models.Page, 0)
	err := p.engine.Find(&datalist)
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
func (p pageService) DeleteByID(id int64) error {
	data := models.Page{Id: id, IsActive: 0}
	_, err := p.engine.Id(data.Id).Update(data)
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

func (p pageService) Update(data *models.Page, column []string) error {
	_, err := p.engine.Id(data.Id).MustCols(column...).Update(data)
	return err
}
func (p pageService) Create(data *models.Page) error {
	_, err := p.engine.Insert(data)
	return err
}
