package models

import (
	"fmt"
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
