package domain

import (
	"github.com/nafisalfiani/ketson-api-gateway-service/lib/log"
)

type Domains struct {
}

func Init(logger log.Interface) *Domains {
	return &Domains{}
}
