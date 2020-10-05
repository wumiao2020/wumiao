package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type NewsService interface {
	GetAll() []models.News
	GetList(limit int, start int) []models.News
	Get(string string) *models.News
	GetByUuid(string string) *models.News
	DeleteByID(id int64) error
	Update(data *models.News, columns []string) error
	Create(data *models.News) error
}

type newsService struct {
	engine *xorm.EngineGroup
}

func NewNewsService() NewsService {
	db := datasource.GetMysqlGroup()
	return &newsService{
		engine: db,
	}
}

func (p newsService) GetList(limit int, start int) []models.News {
	datalist := make([]models.News, 0)
	err := p.engine.Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p newsService) GetAll() []models.News {
	datalist := make([]models.News, 0)
	err := p.engine.Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (p newsService) GetByUuid(string string) *models.News {
	data := &models.News{Uuid: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}
func (p newsService) DeleteByID(id int64) error {
	data := models.News{Id: id, IsActive: 0}
	_, err := p.engine.Id(data.Id).Update(data)
	return err
}
func (p newsService) Get(string string) *models.News {
	data := &models.News{Identifier: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p newsService) Update(data *models.News, column []string) error {
	_, err := p.engine.Where("uuid=?", data.Uuid).MustCols(column...).Update(data)
	return err
}
func (p newsService) Create(data *models.News) error {
	_, err := p.engine.Insert(data)
	return err
}
