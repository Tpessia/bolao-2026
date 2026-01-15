package config

import (
	"strings"
	"testing"
)

var configService ConfigService = ConfigService{}

func TestHealth(t *testing.T) {
	expected := "Healthy"
	result := configService.Health()
	passes := strings.HasPrefix(result, expected)
	if !passes {
		t.Errorf("Health failed with %s", result)
	}
}
