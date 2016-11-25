package gocache

import (
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

type entry struct {
	value interface{}
	time  time.Time
}

type Cache struct {
	m   cmap.ConcurrentMap
	ttl time.Duration
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		m:   cmap.New(),
		ttl: ttl,
	}
}

func (c *Cache) Add(key string, item interface{}) {
	e := &entry{
		value: item,
		time:  time.Now(),
	}
	c.m.Set(key, e)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	t, o := c.m.Get(key)
	if !o {
		return nil, false
	}
	e, _ := t.(*entry)

	if e.time.Add(c.ttl).Before(time.Now()) {
		return nil, false
	}

	return e.value, true
}
