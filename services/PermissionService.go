package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type PermissionService interface {
	GetAll() []models.AdminPermissions
	Get(id int64) *models.AdminPermissions
	Update(data *models.AdminPermissions, columns []string) error
	Create(data *models.AdminPermissions) error
}

type permissionService struct {
	engine *xorm.EngineGroup
}

func NewPermissionService() PermissionService {
	db := datasource.GetMysqlGroup()
	return &permissionService{
		engine: db,
	}
}

func (p permissionService) GetAll() []models.AdminPermissions {
	datalist := make([]models.AdminPermissions, 0)
	err := p.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p permissionService) Get(id int64) *models.AdminPermissions {
	data := &models.AdminPermissions{Id: id}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p permissionService) Update(data *models.AdminPermissions, column []string) error {
	_, err := p.engine.Where("id=?", data.Id).MustCols(column...).Update(data)
	return err
}
func (p permissionService) Create(data *models.AdminPermissions) error {
	_, err := p.engine.Insert(data)
	return err
}
