package config

import (
	"github.com/nafisalfiani/ketson-go-lib/auth"
	"github.com/nafisalfiani/ketson-go-lib/broker"
	"github.com/nafisalfiani/ketson-go-lib/log"
	"github.com/nafisalfiani/ketson-go-lib/security"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/handler/rest"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/handler/scheduler"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/usecase"
)

type Application struct {
	Auth     auth.Config
	Log      log.Config
	Security security.Config
	Broker   broker.Config
	Rest     rest.Config
	Jobs     scheduler.Config
	Usecase  usecase.Config
}
