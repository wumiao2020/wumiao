package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type RoleService interface {
	GetAll() []models.AdminRoles
	Get(id int64) *models.AdminRoles
	Update(data *models.AdminRoles, columns []string) error
	Create(data *models.AdminRoles) error
}

type roleService struct {
	engine *xorm.EngineGroup
}

func NewRoleService() RoleService {
	db := datasource.GetMysqlGroup()
	return &roleService{
		engine: db,
	}
}

func (p roleService) GetAll() []models.AdminRoles {
	datalist := make([]models.AdminRoles, 0)
	err := p.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p roleService) Get(id int64) *models.AdminRoles {
	data := &models.AdminRoles{Id: id}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p roleService) Update(data *models.AdminRoles, column []string) error {
	_, err := p.engine.Where("id=?", data.Id).MustCols(column...).Update(data)
	return err
}
func (p roleService) Create(data *models.AdminRoles) error {
	_, err := p.engine.Insert(data)
	return err
}
