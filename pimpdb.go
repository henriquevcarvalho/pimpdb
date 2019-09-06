package pimpdb

type PimpDB struct {
	Cache Cache
}

func New() *PimpDB {
	p := &PimpDB{}
	return p
}
