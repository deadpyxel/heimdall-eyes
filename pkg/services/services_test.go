package services

import (
	"slices"
	"testing"

	"github.com/deadpyxel/heimdall-eyes/pkg/models"
)

// compare compares two slices of models.Service and returns an integer value indicating their relationship.
// If slice a > b, it returns 1. Is slice a < b, it returns -1.
// If the slices have the same size but any of the fields in the structs differ, it returns -1.
// If the slices are equal, it returns 0.
// Note: This function assumes that the models.Service struct has fields ID, Name, URL, Interval, and SuccessCodes.
func compare(a, b []models.Service) int {
	// If the sizes of the slices differ, return a non zero value
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}
	for i := range a {
		aCurr := a[i]
		bCurr := b[i]
		// Compare all fields of the structs, if any are different, stop the loop and return -1
		// TODO: Find a better way to handle comparison with custom structs that have []int fields
		if aCurr.ID != bCurr.ID || aCurr.Name != bCurr.Name || aCurr.URL != bCurr.URL || aCurr.Interval != bCurr.Interval || slices.Compare(aCurr.SuccessCodes, bCurr.SuccessCodes) != 0 {
			return -1
		}
	}
	// Slices are equal, return "true"
	return 0
}

func TestMonitorManageAddService(t *testing.T) {
	dummySvc := models.Service{ID: 1, Name: "svc", URL: "http://a.b", SuccessCodes: []int{200}}
	svcMngr := ServiceManager{}
	t.Run("When passing an invalid Service returns and error", func(t *testing.T) {
		svc := models.Service{}
		err := svcMngr.AddService(svc)
		if err == nil {
			t.Errorf("Expected error, got nil instead")
		}
	})
	t.Run("When passing an valid Service adds service to slice and returns no error", func(t *testing.T) {
		err := svcMngr.AddService(dummySvc)
		if err != nil {
			t.Errorf("Expected error, got nil instead")
		}
		if len(svcMngr.services) != 1 {
			t.Errorf("Expected services slice to have 1 element, got %v intead", svcMngr.services)
		}
	})

}

func TestMonitorManageRemoveService(t *testing.T) {
	svcMngr := ServiceManager{}
	dummySvc := models.Service{ID: 1, Name: "svc", URL: "http://a.b", SuccessCodes: []int{200}}
	t.Run("When attempting to remove from empty slice returns an error", func(t *testing.T) {
		if err := svcMngr.RemoveService(1); err == nil {
			t.Errorf("Expected an error, got nil instead")
		}
	})
	t.Run("When attempting to remove last existing service returns no error and slice becomes empty", func(t *testing.T) {
		svcMngr.services = []models.Service{dummySvc}
		if err := svcMngr.RemoveService(1); err != nil {
			t.Errorf("Expected no error, got %v instead", err)
		}
		if len(svcMngr.services) != 0 {
			t.Errorf("Expected no remaining services after removal, got %v instead", svcMngr.services)
		}
	})
	t.Run("When attempting to remove non existing service returns an error and no changes are made", func(t *testing.T) {
		contents := []models.Service{dummySvc}
		svcMngr.services = contents
		if err := svcMngr.RemoveService(0); err == nil {
			t.Errorf("Expected an error, got nil instead")
		}
		if compare(svcMngr.services, contents) != 0 {
			t.Errorf("Expected no changes to be made on current services, should be %v, but got %v instead", contents, svcMngr.services)
		}
	})
}
