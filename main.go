package pimpdb

type PimpDB struct {
	Cache *Cache
}

func (p PimpDB) Init() *PimpDB {
	p.Cache = NewCache()
	return &p
}
