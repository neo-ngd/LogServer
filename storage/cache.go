package storage

import (
	"sync"
)

type logcache struct {
	mutex sync.Mutex
	cache map[string]([]LogBody)
	count int
}

func newLogCache(n int) *logcache {
	c := &logcache{
		cache: make(map[string]([]LogBody)),
		count: n,
	}
	return c
}

func (lc *logcache) append(name string, log LogBody) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	logs, ok := lc.cache[name]
	if !ok {
		lc.cache[name] = []LogBody{log}
	}
	logs = append(logs, log)
	if len(logs) <= lc.count {
		lc.cache[name] = logs
	} else {
		lc.cache[name] = logs[1:]
	}
}

func (lc *logcache) getCached(name string) []LogBody {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	result, ok := lc.cache[name]
	if !ok {
		return []LogBody{}
	}
	return result
}

func (lc *logcache) getNames() []string {
	result := []string{}
	for name := range lc.cache {
		result = append(result, name)
	}
	return result
}
