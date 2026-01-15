package config

import "time"

type HealthDetailedRequest struct {
	Input string `json:"input" validate:"required,min=3"`
}

type HealthDetailedResponse struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}
