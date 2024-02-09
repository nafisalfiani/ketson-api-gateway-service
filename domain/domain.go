package domain

import (
	"github.com/nafisalfiani/ketson-go-lib/log"
)

type Domains struct {
}

func Init(logger log.Interface) *Domains {
	return &Domains{}
}
