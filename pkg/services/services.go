package services

import (
	"fmt"

	"github.com/deadpyxel/heimdall-eyes/pkg/models"
)

type MonitorController interface {
	AddService(svc models.Service) error
	RemoveService(id int) error
	GetService(id int) (models.Service, error)
	GetAllServices() []models.Service
}

type ServiceManager struct {
	services []models.Service
}

func (sm *ServiceManager) AddService(svc models.Service) error {
	if err := svc.Validate(); err != nil {
		return err
	}
	sm.services = append(sm.services, svc)
	return nil
}

func (sm *ServiceManager) RemoveService(id int) error {
	for i, svc := range sm.services {
		if svc.ID == id {
			sm.services = append(sm.services[:i], sm.services[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Service with ID %d not found", id)
}
