package managers

import (
	"arc/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

type TableManager struct {
	db *gorm.DB
}

func NewTableManager(db *gorm.DB) *TableManager {
	return &TableManager{
		db: db,
	}
}

func (m *TableManager) GetSolvedProblems(parameter *model.GetSolvedProblemsParameter) ([]*model.SolvedProblem, error) {
	var problems []*model.SolvedProblem
	whereSituation := model.SolvedProblem{
		MainBody: parameter.MainBody,
		Solution: parameter.Solution,
	}
	if err := m.db.Where(whereSituation).Find(&problems).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		logs.Error(err)
		return nil, err
	}
	return problems, nil
}

func (m *TableManager) GetPurposes(parameter *model.GetSolvedProblemsParameter) ([]string, error) {
	units, err := m.GetSolvedProblems(parameter)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var purpose []string
	for _, unit := range units {
		purpose = append(purpose, unit.Purpose)
	}
	return purpose, nil
}

func (m *TableManager) InsertIfNotExist(tables ...model.Table) error {
	for _, table := range tables {
		if err := m.createTableIfNotExist(table); err != nil {
			logs.Error(err)
			return err
		}
		isExist, err := m.isExist(table)
		if err != nil {
			logs.Error(err)
			return err
		}
		if isExist {
			logs.Info("%s'%+v has exist, so don't need to insert.", table.TableName(), table.GetIdentity())
			continue
		}
		if err := m.insert(table); err != nil {
			logs.Error(err)
			return err
		}

	}
	return nil
}

func (m *TableManager) isExist(table model.Table) (bool, error) {
	if !m.db.HasTable(table) {
		return false, nil
	}
	if err := m.db.Where(table.GetIdentity()).First(table).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		logs.Error(err)
		return false, err
	}
	return true, nil
}

func (m *TableManager) createTableIfNotExist(table model.Table) error {
	if !m.db.HasTable(table) {
		if err := m.db.Error; err != nil {
			logs.Error(err)
			return err
		}
		if err := m.db.CreateTable(table).Error; err != nil {
			logs.Error(err)
			return err
		}
	}
	return nil
}

func (m *TableManager) insert(tables ...model.Table) error {
	for _, table := range tables {
		if err := m.createTableIfNotExist(table); err != nil {
			logs.Error(err)
			return err
		}
		if err := m.db.Create(table).Error; err != nil {
			logs.Error(err)
			return err
		}
	}
	return nil
}
