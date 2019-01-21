package api

import (
	"sync"
)

type logcache struct {
	mutex sync.Mutex
	cache map[string]([]logBody)
	count int
}

func NewLogCache(n int) *logcache {
	c := &logcache{
		cache: make(map[string]([]logBody)),
		count: n,
	}
	return c
}

func (lc *logcache) append(name string, log logBody) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	logs, ok := lc.cache[name]
	if !ok {
		lc.cache[name] = []logBody{log}
	}
	logs = append(logs, log)
	if len(logs) <= lc.count {
		lc.cache[name] = logs
	} else {
		lc.cache[name] = logs[1:]
	}
}
func (lc *logcache) Append(name string, log logBody) {
	lc.append(name, log)
	lc.append("all", log)
}

func (lc *logcache) GetCached(name string) []logBody {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	result, ok := lc.cache[name]
	if !ok {
		return []logBody{}
	}
	return result
}
