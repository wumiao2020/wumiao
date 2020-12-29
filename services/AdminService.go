package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type AdminService interface {
	GetAll() []models.Admins
	GetList(limit int, start int) []models.Admins
	GetById(Id int64) *models.Admins
	GetByEmail(email string) *models.Admins
	DeleteByID(id int64) error
	Update(data *models.Admins, columns []string) error
	Create(data *models.Admins) error
}

type adminService struct {
	engine *xorm.EngineGroup
}

func NewAdminService() AdminService {
	db := datasource.GetMysqlGroup()
	return &adminService{
		engine: db,
	}
}

func (a adminService) GetList(limit int, start int) []models.Admins {
	datalist := make([]models.Admins, 0)
	err := a.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (a adminService) GetAll() []models.Admins {
	datalist := make([]models.Admins, 0)
	err := a.engine.Where("status=?", 0).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (a adminService) GetById(id int64) *models.Admins {
	data := &models.Admins{Id: id}
	ok, err := a.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (a adminService) GetByEmail(email string) *models.Admins {
	data := &models.Admins{Email: email}
	ok, err := a.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (a adminService) DeleteByID(id int64) error {
	data := models.Admins{Id: id, Status: 0}
	_, err := a.engine.Id(data.Id).Update(data)
	return err
}

func (a adminService) Update(data *models.Admins, column []string) error {
	_, err := a.engine.Id(data.Id).MustCols(column...).Update(data)
	return err
}
func (a adminService) Create(data *models.Admins) error {
	_, err := a.engine.Insert(data)
	return err
}
