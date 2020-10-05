package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type MenuService interface {
	GetAll() []models.Menu
	GetList(limit int, start int) []models.Menu
	GetById(id int64) *models.Menu
	DeleteByID(id int64) error
	Update(data *models.Menu, columns []string) error
	Create(data *models.Menu) error
}

type menuService struct {
	engine *xorm.EngineGroup
}

func NewMenuService() MenuService {
	db := datasource.GetMysqlGroup()
	return &menuService{
		engine: db,
	}
}

func (p menuService) GetList(limit int, start int) []models.Menu {
	datalist := make([]models.Menu, 0)
	err := p.engine.Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p menuService) GetAll() []models.Menu {
	datalist := make([]models.Menu, 0)
	err := p.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (p menuService) GetById(id int64) *models.Menu {
	data := &models.Menu{Id: id}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}
func (p menuService) DeleteByID(id int64) error {
	data := models.Menu{Id: id, IsActive: 0}
	_, err := p.engine.Id(data.Id).Update(data)
	return err
}

func (p menuService) Update(data *models.Menu, column []string) error {
	_, err := p.engine.Where("id=?", data.Id).MustCols(column...).Update(data)
	return err
}
func (p menuService) Create(data *models.Menu) error {
	_, err := p.engine.Insert(data)
	return err
}
