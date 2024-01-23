package main

import (
	"fmt"

	"github.com/deadpyxel/heimdall-eyes/pkg/monitor"
	"github.com/deadpyxel/heimdall-eyes/pkg/services"
)

func main() {
	svcManager := services.ServiceManager{}
	httpMonitor := monitor.HTTPMonitor{}

	fmt.Printf("%v\n%v", svcManager, httpMonitor)
}
