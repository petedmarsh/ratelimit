package memcached

import (
	"github.com/rainycape/memcache"
)

var _ Client = (*memcache.Client)(nil)

// Interface for memcached, used for mocking.
type Client interface {
	GetMulti(keys []string) (map[string]*memcache.Item, error)
	Increment(key string, delta uint64) (newValue uint64, err error)
	Add(item *memcache.Item) error
}
