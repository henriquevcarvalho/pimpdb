package pimpdb

import (
	"log"
)

type PimpDB struct {
	Cache 		*Cache
	Log			*log.Logger
}

func New() *PimpDB {
	return &PimpDB{}
}