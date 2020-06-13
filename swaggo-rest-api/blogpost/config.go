package blogpost

import (
	"github.com/bygui86/go-swagger/swaggo-rest-api/logging"
	"github.com/bygui86/go-swagger/swaggo-rest-api/utils"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostEnvVarDefault = "localhost"
	restPortEnvVarDefault = 8080
)

func loadConfig() *config {
	logging.Log.Debug("Load configurations")
	return &config{
		RestHost: utils.GetStringEnv(restHostEnvVar, restHostEnvVarDefault),
		RestPort: utils.GetIntEnv(restPortEnvVar, restPortEnvVarDefault),
	}
}
