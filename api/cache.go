package api

import (
	"sync"
)

type logcache struct {
	mutex sync.Mutex
	cache []logBody
	count int
}

func NewLogCache(n int) *logcache {
	c := &logcache{
		cache: []logBody{},
		count: n,
	}
	return c
}

func (lc *logcache) Append(log logBody) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	tmp := append(lc.cache, log)
	if len(tmp) <= lc.count {
		lc.cache = tmp
	} else {
		lc.cache = tmp[1:]
	}
}

func (lc *logcache) GetCached() []logBody {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	result := []logBody{}
	for _, v := range lc.cache {
		result = append(result, v)
	}
	return result
}
