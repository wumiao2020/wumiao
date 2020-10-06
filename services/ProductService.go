package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type ProductService interface {
	GetAll() []models.Product
	GetList(limit int, start int) []models.Product
	GetCategoryList(parentId int, limit int, start int) []models.Product
	GetTopList(limit int, start int) []models.Product
	Get(string string) *models.Product
	GetByUuid(string string) *models.Product
	DeleteByID(id int64) error
	Update(data *models.Product, columns []string) error
	Create(data *models.Product) error
}

type productService struct {
	engine *xorm.EngineGroup
}

func NewProductService() ProductService {
	db := datasource.GetMysqlGroup()
	return &productService{
		engine: db,
	}
}

func (p productService) GetCategoryList(parentId int, limit int, start int) []models.Product {
	datalist := make([]models.Product, 0)
	err := p.engine.Where("parent_id=?", parentId).Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p productService) GetList(limit int, start int) []models.Product {
	datalist := make([]models.Product, 0)
	err := p.engine.Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p productService) GetTopList(limit int, start int) []models.Product {
	datalist := make([]models.Product, 0)
	err := p.engine.Desc("position", "id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p productService) GetAll() []models.Product {
	datalist := make([]models.Product, 0)
	err := p.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
func (p productService) GetByUuid(string string) *models.Product {
	data := &models.Product{Uuid: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}
func (p productService) DeleteByID(id int64) error {
	data := models.Product{Id: id, IsActive: 0}
	_, err := p.engine.Id(data.Id).Update(data)
	return err
}
func (p productService) Get(string string) *models.Product {
	data := &models.Product{Identifier: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p productService) Update(data *models.Product, column []string) error {
	_, err := p.engine.Where("uuid=?", data.Uuid).MustCols(column...).Update(data)
	return err
}
func (p productService) Create(data *models.Product) error {
	_, err := p.engine.Insert(data)
	return err
}
