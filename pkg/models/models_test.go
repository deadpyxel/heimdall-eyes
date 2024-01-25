package models

import (
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("When Service has no ID Validate returns and error", func(t *testing.T) {
		svc := Service{}
		err := svc.Validate()
		exp_err := "Service must have an ID"
		if err == nil {
			t.Errorf("Expected an error, got nil instead")
		}
		if err.Error() != exp_err {
			t.Errorf("Expected error to be [%s], got [%s] instead", exp_err, err.Error())
		}
	})
	t.Run("When Service has no name Validate returns and error", func(t *testing.T) {
		svc := Service{ID: 1}
		err := svc.Validate()
		exp_err := "Service name is required"
		if err == nil {
			t.Errorf("Expected an error, got nil instead")
		}
		if err.Error() != exp_err {
			t.Errorf("Expected error to be [%s], got [%s] instead", exp_err, err.Error())
		}
	})
	t.Run("When Service has empty URL Validate returns and error", func(t *testing.T) {
		svc := Service{ID: 1, Name: "SVC"}
		err := svc.Validate()
		exp_err := "empty url"
		if err == nil {
			t.Errorf("Expected an error, got nil instead")
		}
		if !strings.Contains(err.Error(), exp_err) {
			t.Errorf("Expected error to contain [%s], got [%s] instead", exp_err, err.Error())
		}
	})
	t.Run("When Service has invalid URL Validate returns and error", func(t *testing.T) {
		svc := Service{ID: 1, Name: "SVC", URL: "url"}
		err := svc.Validate()
		exp_err := "invalid URI for request"
		if err == nil {
			t.Errorf("Expected an error, got nil instead")
		}
		if !strings.Contains(err.Error(), exp_err) {
			t.Errorf("Expected error to contain [%s], got [%s] instead", exp_err, err.Error())
		}
	})
	t.Run("When Service has no success codes Validate returns and error", func(t *testing.T) {
		svc := Service{ID: 1, Name: "SVC", URL: "http://a.b"}
		err := svc.Validate()
		exp_err := "At least one success code is required"
		if err == nil {
			t.Errorf("Expected an error, got nil instead")
		}
		if !strings.Contains(err.Error(), exp_err) {
			t.Errorf("Expected error to contain [%s], got [%s] instead", exp_err, err.Error())
		}
	})
	t.Run("When Service has all required fields Validate returns no error", func(t *testing.T) {
		svc := Service{ID: 1, Name: "SVC", URL: "http://a.b", SuccessCodes: []int{200}}
		err := svc.Validate()
		if err != nil {
			t.Errorf("Expected no error, got %v instead", err)
		}
	})
}
