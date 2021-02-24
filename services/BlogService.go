package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type BlogService interface {
	GetAll() []models.Blog
	GetList(limit int, start int) []models.Blog
	Get(string string) *models.Blog
	GetByUuid(string string) *models.Blog
	DeleteByID(id int64) error
	Update(data *models.Blog, columns []string) error
	Create(data *models.Blog) error
}

type blogService struct {
	engine *xorm.EngineGroup
}

func NewBlogService() BlogService {
	db := datasource.GetMysqlGroup()
	return &blogService{
		engine: db,
	}
}

func (p blogService) GetList(limit int, start int) []models.Blog {
	datalist := make([]models.Blog, 0)
	err := p.engine.Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p blogService) GetAll() []models.Blog {
	datalist := make([]models.Blog, 0)
	err := p.engine.Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (p blogService) GetByUuid(string string) *models.Blog {
	data := &models.Blog{Uuid: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}
func (p blogService) DeleteByID(id int64) error {
	data := models.Blog{Id: id, Status: 0}
	_, err := p.engine.Id(data.Id).Update(data)
	return err
}
func (p blogService) Get(string string) *models.Blog {
	data := &models.Blog{Identifier: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p blogService) Update(data *models.Blog, column []string) error {
	_, err := p.engine.Where("uuid=?", data.Uuid).MustCols(column...).Update(data)
	return err
}
func (p blogService) Create(data *models.Blog) error {
	_, err := p.engine.Insert(data)
	return err
}
