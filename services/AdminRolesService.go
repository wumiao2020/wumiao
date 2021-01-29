package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type AdminRolesService interface {
	GetAll() []models.AdminRoles
	GetList(limit int, start int) []models.AdminRoles
	GetById(Id int64) *models.AdminRoles
	DeleteByID(id int64) error
	Update(data *models.AdminRoles) error
	Create(data *models.AdminRoles) error
}

type adminRolesService struct {
	engine *xorm.EngineGroup
}

func NewAdminRolesService() AdminRolesService {
	db := datasource.GetMysqlGroup()
	return &adminRolesService{
		engine: db,
	}
}

func (a adminRolesService) GetList(limit int, start int) []models.AdminRoles {
	datalist := make([]models.AdminRoles, 0)
	err := a.engine.Desc("id").Limit(limit, start).Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (a adminRolesService) GetAll() []models.AdminRoles {
	datalist := make([]models.AdminRoles, 0)
	err := a.engine.Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (a adminRolesService) GetById(id int64) *models.AdminRoles {
	data := &models.AdminRoles{Id: id}
	ok, err := a.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (a adminRolesService) DeleteByID(id int64) error {
	data := models.AdminRoles{Id: id, Status: 0}
	_, err := a.engine.Id(data.Id).Delete(data)
	return err
}

func (a adminRolesService) Update(data *models.AdminRoles) error {
	column := []string{"status"}
	_, err := a.engine.Id(data.Id).MustCols(column...).Update(data)
	return err
}
func (a adminRolesService) Create(data *models.AdminRoles) error {
	_, err := a.engine.Insert(data)
	return err
}
