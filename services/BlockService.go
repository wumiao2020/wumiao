package services

import (
	"github.com/go-xorm/xorm"
	"wumiao/datasource"
	"wumiao/models"
)

type BlockService interface {
	GetAll(search string) []models.Block
	GetList(limit int, start int, search string) []models.Block
	Get(string string) *models.Block
	GetById(id int64) *models.Block
	GetByIdentifier(identifier string) *models.Block
	DeleteByID(id int64) error
	Update(data *models.Block) error
	Create(data *models.Block) error
}

type blockService struct {
	engine *xorm.EngineGroup
}

func NewBlockService() BlockService {
	db := datasource.GetMysqlGroup()
	return &blockService{
		engine: db,
	}
}

func (p blockService) GetList(limit int, start int, search string) []models.Block {
	datalist := make([]models.Block, 0)
	var err error
	if len(search) > 0 {
		err = p.engine.Desc("id").Where("title like ?", "%"+search+"%").Limit(limit, start).Find(&datalist)
	} else {
		err = p.engine.Desc("id").Limit(limit, start).Find(&datalist)
	}
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p blockService) GetAll(search string) []models.Block {
	datalist := make([]models.Block, 0)
	var err error
	if len(search) > 0 {
		err = p.engine.Desc("id").Where("title like ?", "%"+search+"%").Find(&datalist)
	} else {
		err = p.engine.Desc("id").Find(&datalist)
	}
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (p blockService) GetById(id int64) *models.Block {
	data := &models.Block{Id: id}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p blockService) GetByIdentifier(identifier string) *models.Block {
	data := &models.Block{Identifier: identifier}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return data
	}
}
func (p blockService) DeleteByID(id int64) error {
	data := models.Block{Id: id}
	_, err := p.engine.Id(data.Id).Delete(data)
	return err
}
func (p blockService) Get(string string) *models.Block {
	data := &models.Block{Identifier: string}
	ok, err := p.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
	}
}

func (p blockService) Update(data *models.Block) error {
	_, err := p.engine.Where("id=?", data.Id).AllCols().Update(data)
	return err
}
func (p blockService) Create(data *models.Block) error {
	_, err := p.engine.Insert(data)
	return err
}
