package pimpdb

import (
	"log"
)

type PimpDB struct {
	Cache 		*Cache
	Log			*log.Logger
}

func (p PimpDB) Init() *PimpDB {
	p.Cache = NewCache()
	return &p
}