package services

import (
	"github.com/htquangg/a-wasm/config"
	"github.com/htquangg/a-wasm/internal/protocluster"
	"github.com/htquangg/a-wasm/internal/repos"
	"github.com/htquangg/a-wasm/internal/services/api_key"
	"github.com/htquangg/a-wasm/internal/services/auth"
	"github.com/htquangg/a-wasm/internal/services/deployment"
	"github.com/htquangg/a-wasm/internal/services/endpoint"
	"github.com/htquangg/a-wasm/internal/services/health"
	"github.com/htquangg/a-wasm/internal/services/mailer"
	"github.com/htquangg/a-wasm/internal/services/session"
	"github.com/htquangg/a-wasm/internal/services/user"
)

type Sevices struct {
	cfg          *config.Config
	repos        *repos.Repos
	protocluster *protocluster.Cluster
	Health       *health.HealthService
	Mailer       *mailer.MailerService
	Endpoint     *endpoint.EndpointService
	Deployment   *deployment.DeploymentService
	Auth         *auth.AuthService
	Session      *session.SessionService
	User         *user.UserService
	ApiKey       *api_key.ApiKeyService
}

func New(cfg *config.Config, repos *repos.Repos, protoCluster *protocluster.Cluster) *Sevices {
	healthService := health.NewHealthService(repos.Health)
	mailerService := mailer.NewMailerService(cfg, repos.Mailer)
	endpointService := endpoint.NewEndpointService(
		cfg,
		repos.Endpoint,
		repos.DeploymentCommon,
		protoCluster,
	)
	deploymentService := deployment.NewDeploymentService(
		cfg,
		repos.Deployment,
		repos.EndpointCommon,
		protoCluster,
	)
	authService := auth.NewAuthService(repos.Auth)
	sessionService := session.NewSessionService(cfg, repos.Session, repos.UserCommon)
	userService := user.NewUserService(
		cfg,
		repos.User,
		repos.UserAuth,
		sessionService,
		mailerService,
	)
	apiKeyService := api_key.NewApiKeyService(cfg, repos.ApiKey)

	return &Sevices{
		cfg:          cfg,
		repos:        repos,
		protocluster: protoCluster,
		Health:       healthService,
		Endpoint:     endpointService,
		Deployment:   deploymentService,
		Auth:         authService,
		Session:      sessionService,
		User:         userService,
		ApiKey:       apiKeyService,
	}
}
