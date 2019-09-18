package pimpdb

import (
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
)

type Cache struct {
	Service *cache.Cache
}

func (p *PimpDB) SetCacheOptions(opt ...Cache) {
	if len(opt) == 0 {
		p.Cache.Service = cache.New(cache.NoExpiration, cache.NoExpiration)
	} else {
		p.Cache = opt[0]
	}
}

func (p *Cache) checkSet(id string, x interface{}) error {
	LogDefault(id, x, "check_set")
	return p.Service.Add(id, x, cache.NoExpiration)
}

func (p *Cache) Set(id string, x interface{}, exists bool) error {
	if exists {
		return p.checkSet(id, x)
	}
	LogDefault(id, x, "set")
	p.Service.Set(id, x, cache.NoExpiration)
	return errors.New("")
}

func (p *Cache) Replace(id string, x interface{}) error {
	LogDefault(id, x, "replace")
	return p.Service.Replace(id, x, cache.NoExpiration)
}

func (p *Cache) Get(id string) (interface{}, bool) {
	LogDefault(id, nil,"get")
	val, found := p.Service.Get(id)
	return val, found
}
