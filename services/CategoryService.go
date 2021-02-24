package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type CategoryService interface {
	GetAll() []models.Category
	GetList(limit int, start int) []models.Category
	Get(string string) *models.Category
	GetByIdentifier(string string) *models.Category
	GetByUuid(string string) *models.Category
	DeleteByID(id int64) error
	Update(data *models.Category, columns []string) error
	Create(data *models.Category) error
}

type categoryService struct {
	engine *xorm.EngineGroup
}

func NewCategoryService() CategoryService {
	db := datasource.GetMysqlGroup()
	return &categoryService{
		engine: db,
	}
}

func (p categoryService) GetList(limit int, start int) []models.Category {
	datalist := make([]models.Category, 0)
	err := p.engine.Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p categoryService) GetAll() []models.Category {
	datalist := make([]models.Category, 0)
	err := p.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (p categoryService) GetByUuid(string string) *models.Category {
	data := &models.Category{Uuid: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p categoryService) GetByIdentifier(string string) *models.Category {
	data := &models.Category{Identifier: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}
func (p categoryService) DeleteByID(id int64) error {
	data := models.Category{Id: id, Status: 0}
	_, err := p.engine.Id(data.Id).Update(data)
	return err
}
func (p categoryService) Get(string string) *models.Category {
	data := &models.Category{Identifier: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p categoryService) Update(data *models.Category, column []string) error {
	_, err := p.engine.Where("uuid=?", data.Uuid).MustCols(column...).Update(data)
	return err
}
func (p categoryService) Create(data *models.Category) error {
	_, err := p.engine.Insert(data)
	return err
}
