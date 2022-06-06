package service

import (
	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"gorm.io/gorm"
)

// search total unique user agent

type serviceTracker struct {
	db *gorm.DB
}

func NewTracker(db *gorm.DB) domain.TrackerUseCase {
	return &serviceTracker{
		db: db,
	}
}
func (s serviceTracker) SearchUniqueUserAgent(endpoint string) (int, error) {
	var data []domain.CounterEndpoint
	err := s.db.Model(&domain.CounterEndpoint{}).Where("end_point = ?", endpoint).Find(&data).Error
	if data == nil && err != nil {
		return 0, err
	} else if data == nil && err == nil {
		return 0, nil
	}
	return len(data), nil
}

func (s serviceTracker) SaveTracker(id uint, endpoint string) error {
	//search if endpoint already exist
	var data domain.CounterEndpoint
	err := s.db.Model(&domain.CounterEndpoint{}).Where("unique_user_agent = ?", id).Where("end_point = ?", endpoint).Take(&data).Error
	if data == (domain.CounterEndpoint{}) {
		var newEndpoint domain.CounterEndpoint
		newEndpoint.UniqueUserAgent = int(id)
		newEndpoint.Count = 1
		newEndpoint.EndPoint = endpoint
		if err := s.db.Model(&domain.CounterEndpoint{}).Create(&newEndpoint).Error; err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	if data.EndPoint == endpoint {
		temp := data.Count + 1
		if err := s.db.Model(&domain.CounterEndpoint{}).Where("unique_user_agent = ?", id).Update("count", temp).Error; err != nil {
			return err
		}
		return nil
	}
	if data.EndPoint != endpoint {
		var newEndpoint domain.CounterEndpoint
		newEndpoint.UniqueUserAgent = int(id)
		newEndpoint.Count = 1
		newEndpoint.EndPoint = endpoint
		if err := s.db.Model(&domain.CounterEndpoint{}).Create(&newEndpoint).Error; err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s serviceTracker) SearchEndpointCalled(endpoint string) (int, error) {
	var data []domain.CounterEndpoint
	err := s.db.Model(&domain.CounterEndpoint{}).Where("end_point = ?", endpoint).Find(&data).Error
	if data == nil && err != nil {
		return 0, err
	} else if data == nil && err == nil {
		return 0, nil
	}
	total := 0
	for _, v := range data {
		total += v.Count
	}
	return total, nil
}
