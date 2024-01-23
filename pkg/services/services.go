package services

import "github.com/deadpyxel/heimdall-eyes/pkg/models"

type MonitorController interface {
	AddService(svc models.Service) error
	RemoveService(id int) error
	GetService(id int) (models.Service, error)
	GetAllServices() []models.Service
}

type ServiceManager struct {
	services []models.Service
}
