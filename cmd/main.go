package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/deadpyxel/heimdall-eyes/pkg/models"
	"github.com/deadpyxel/heimdall-eyes/pkg/monitor"
	"github.com/deadpyxel/heimdall-eyes/pkg/services"
)

func main() {
	svcManager := services.ServiceManager{}
	httpMonitor := monitor.HTTPMonitor{}

	// Example service
	service := models.Service{
		ID:           1,
		Name:         "Example Service",
		URL:          "https://github.com",
		SuccessCodes: []int{200},
		Interval:     30 * time.Second,
	}

	// Add service to the manager
	if err := svcManager.AddService(service); err != nil {
		fmt.Printf("Error adding service: %v\n", err)
		return
	}

	done := make(chan struct{})
	go httpMonitor.StartMonitoring(service, done)

	// Graceful shutdown handling
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	fmt.Println("Shutting down...")
	close(done)
}
