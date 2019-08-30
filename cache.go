package pimpdb

import (
	"fmt"
	"github.com/patrickmn/go-cache"
)

type Cache struct {
	Service *cache.Cache
	*PimpDB
}

type CachedSession struct {
	User      string
	SessionId string
	*Cache
}

func NewCache() *Cache {
	c := new(Cache)
	c.Service = cache.New(cache.NoExpiration, cache.NoExpiration)
	return c
}

func (p *PimpDB) Save(id string, x interface{}) error {
	fmt.Println("[x] Pimping : "+id, x)
	return p.Cache.Service.Add(id, x, cache.NoExpiration)
}

func (p *PimpDB) Get(id string) (interface{}, bool) {
	val, found := p.Cache.Service.Get(id)
	return val, found
}
