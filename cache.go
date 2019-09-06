package pimpdb

import (
	"github.com/patrickmn/go-cache"
)

type Cache struct {
	Service 		*cache.Cache
}

type CachedSession struct {
	User      		string
	SessionId 		string
	*Cache
}

func (p *PimpDB) SetCacheOptions(opt ...Cache) {
	if len(opt) == 0 {
		p.Cache.Service = cache.New(cache.NoExpiration, cache.NoExpiration)
	} else {
		p.Cache = opt[0]
	}
}

func (p *Cache) Save(id string, x interface{}) error {
	LogSave(id, x)
	return p.Service.Add(id, x, cache.NoExpiration)
}

func (p *Cache) Get(id string) (interface{}, bool) {
	LoveGet(id)
	val, found := p.Service.Get(id)
	return val, found
}
