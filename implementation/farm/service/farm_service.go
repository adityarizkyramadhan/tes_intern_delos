package service

import (
	"fmt"
	"log"

	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"gorm.io/gorm"
)

type serviceFarm struct {
	db *gorm.DB
	tr domain.TrackerUseCase
}

func NewServiceFarm(db *gorm.DB, tracker domain.TrackerUseCase) domain.FarmUseCase {
	return &serviceFarm{
		db: db,
		tr: tracker,
	}
}

func (s serviceFarm) Create(model *domain.FarmModel) (uint, error) {
	if err := s.db.Model(domain.FarmModel{}).Create(&model).Error; err != nil {
		return 0, err
	}
	if err := s.tr.SaveTracker(model.ID, "/farm"); err != nil {
		return 0, err
	}
	fmt.Println("id: ", model.ID)
	return model.ID, nil
}

func (s serviceFarm) Read(id uint) (domain.FarmModel, error) {
	if err := s.tr.SaveTracker(id, "/farm"); err != nil {
		return domain.FarmModel{}, err
	}
	var readFarm domain.FarmModel
	if err := s.db.Model(&domain.FarmModel{}).Where("id = ?", id).Preload("PondModels").Find(&readFarm).Error; err != nil {
		return domain.FarmModel{}, err
	}
	return readFarm, nil
}

func (s serviceFarm) Delete(id uint) error {
	if err := s.tr.SaveTracker(id, "/farm"); err != nil {
		return err
	}
	if err := s.db.Model(&domain.FarmModel{}).Unscoped().Delete("id = ?", id).Error; err != nil {
		return err
	}
	if err := s.db.Model(&domain.PondModel{}).Unscoped().Delete("farm_entity_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (s serviceFarm) Update(model *domain.FarmModel, id uint) error {
	if err := s.tr.SaveTracker(id, "/farm"); err != nil {
		log.Println(err)
		return err
	}
	if err := s.db.Model(&domain.FarmModel{}).Where("id = ?", id).Updates(&model).Error; err != nil {
		return err
	}
	return nil
}

func (s serviceFarm) ReadAll() ([]domain.FarmModel, error) {
	var data []domain.FarmModel
	if err := s.db.Model(&domain.FarmModel{}).Preload("PondModels").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
