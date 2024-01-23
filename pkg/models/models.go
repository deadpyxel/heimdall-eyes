package models

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

type Service struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	URL          string        `json:"url"`
	SuccessCodes []int         `json:"success_codes"`
	Interval     time.Duration `json:"interval"`
}

func (s *Service) String() string {
	return fmt.Sprintf("Service: %s [%s]", s.Name, s.URL)
}

func (s *Service) Validate() error {
	if s.Name == "" {
		return errors.New("Service name is required")
	}
	if _, err := url.ParseRequestURI(s.URL); err != nil {
		return fmt.Errorf("Invalid URL pattern: %v", err)
	}
	if len(s.SuccessCodes) == 0 {
		return errors.New("At least one success code is required")
	}
	return nil
}
