package monitor

import (
	"fmt"

	"github.com/deadpyxel/heimdall-eyes/pkg/models"
)

type Monitor interface {
	StartMonitoring(svc models.Service)
}

type HTTPMonitor struct{}

func (m *HTTPMonitor) StartMonitoring(svc models.Service) {
	fmt.Printf("Started monitoring uptime of %v", svc)
}
