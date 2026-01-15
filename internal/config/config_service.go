package config

import (
	"fmt"
	"time"
)

type ConfigService struct{}

func (s *ConfigService) Health() string {
	healthMsg := fmt.Sprintf("Healthy! %v", time.Now())
	return healthMsg
}

func (s *ConfigService) HealthDetailed() HealthDetailed {
	healthDetailed := HealthDetailed{
		Status: "Healthy",
		Time:   time.Now(),
	}

	return healthDetailed
}
