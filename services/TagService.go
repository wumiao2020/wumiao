package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type TagService interface {
	GetAll() []models.Tag
	GetList(limit int, start int) []models.Tag
	Get(string string) *models.Tag
	GetByUuid(string string) *models.Tag
	DeleteByID(id int64) error
	Update(data *models.Tag, columns []string) error
	Create(data *models.Tag) error
}

type tagService struct {
	engine *xorm.EngineGroup
}

func NewTagService() TagService {
	db := datasource.GetMysqlGroup()
	return &tagService{
		engine: db,
	}
}

func (p tagService) GetList(limit int, start int) []models.Tag {
	datalist := make([]models.Tag, 0)
	err := p.engine.Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p tagService) GetAll() []models.Tag {
	datalist := make([]models.Tag, 0)
	err := p.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (p tagService) GetByUuid(string string) *models.Tag {
	data := &models.Tag{Uuid: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}
func (p tagService) DeleteByID(id int64) error {
	data := models.Tag{Id: id, IsActive: 0}
	_, err := p.engine.Id(data.Id).Update(data)
	return err
}
func (p tagService) Get(string string) *models.Tag {
	data := &models.Tag{Identifier: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p tagService) Update(data *models.Tag, column []string) error {
	_, err := p.engine.Where("uuid=?", data.Uuid).MustCols(column...).Update(data)
	return err
}
func (p tagService) Create(data *models.Tag) error {
	_, err := p.engine.Insert(data)
	return err
}
