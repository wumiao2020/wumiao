package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type PermissionService interface {
	GetAll() []models.AdminPermissions
	GetMenuList() []models.AdminPermissions
	Get(id int64) *models.AdminPermissions
	GetBreadcrumbs(url string) []models.AdminPermissions
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
	err := p.engine.Asc("id").Find(&datalist)
	if err != nil {
		println(err)
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

func (p permissionService) get(id int64) *models.AdminPermissions {
	data := &models.AdminPermissions{Id: id}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p permissionService) GetBreadcrumbs(url string) []models.AdminPermissions {
	data := &models.AdminPermissions{Perms: url}
	ok, err := p.engine.Get(data)
	breadcrumbs := recursionBreadcrumbs(p.GetAll(), data.Id)
	if ok && err == nil {
		return breadcrumbs
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

func (p permissionService) GetMenuList() []models.AdminPermissions {
	return recursionMenuList(p.GetAll(), 0, 1)
}

//递归函数
func recursionMenuList(data []models.AdminPermissions, pid int64, level int64) []models.AdminPermissions {
	var listTree []models.AdminPermissions
	for _, value := range data {
		if value.ParentId == pid {
			value.Children = recursionMenuList(data, value.Id, level+1)
			listTree = append(listTree, value)
		}
	}
	return listTree
}

//递归函数
func recursionBreadcrumbs(data []models.AdminPermissions, pid int64) []models.AdminPermissions {
	var listTree []models.AdminPermissions
	for _, value := range data {
		if value.Id == pid && pid != 0 {
			vars := recursionBreadcrumbs(data, value.ParentId)
			listTree = append(vars, value)
		}
	}
	return listTree
}
