package controllers

import "github.com/htquangg/a-wasm/internal/services"

type Controllers struct {
	Health     *HealthController
	Endpoint   *EndpointController
	Deployment *DeploymentController
}

func New(services *services.Sevices) *Controllers {
	return &Controllers{
		Health:     NewHealthController(services.Health),
		Endpoint:   NewEndpointController(services.Endpoint),
		Deployment: NewDeploymentController(services.Deployment),
	}
}