package services

import "hexagonal-example/internal/core/domain/health"

type HealthService struct {
	Health health.Health
}

func (healthSrv HealthService) IsAppHealthy() bool {
	status := healthSrv.Health.CheckHealth()
	return status
}
