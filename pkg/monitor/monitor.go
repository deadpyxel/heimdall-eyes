package monitor

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deadpyxel/heimdall-eyes/pkg/models"
)

type Monitor interface {
	StartMonitoring(svc models.Service, done chan<- struct{})
}

type HTTPMonitor struct{}

func (m *HTTPMonitor) StartMonitoring(svc models.Service, done <-chan struct{}) {
	ticker := time.NewTicker(svc.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Perform the monitoring check
			resp, err := http.Get(svc.URL)
			if err != nil {
				fmt.Printf("Error monitoring service %s: %v\n", svc.Name, err)
				continue
			}
			defer resp.Body.Close()

			if m.isResponseSuccessful(svc.SuccessCodes, resp.StatusCode) {
				fmt.Printf("Service %s is up. Response code: %d\n", svc.Name, resp.StatusCode)
			} else {
				fmt.Printf("Service %s is down. Response code: %d\n", svc.Name, resp.StatusCode)
			}
		case <-done:
			fmt.Printf("Stopped monitoring service: %s\n", svc.Name)
			return
		}
	}
}

func (m *HTTPMonitor) isResponseSuccessful(successCodes []int, responseCode int) bool {
	for _, code := range successCodes {
		if code == responseCode {
			return true
		}
	}
	return false
}
