package storage

import (
	"fmt"
	"time"
)

type LogBody struct {
	Text string
	Name string
	Key  int64
}

type Storage struct {
	p        *persist
	c        *logcache
	registee []chan<- LogBody
}

const cache_count = 50

func NewStorage(path, name string, age, split time.Duration) *Storage {
	return &Storage{
		p: newPersist(path, name, age, split),
		c: newLogCache(cache_count),
	}
}
func (s *Storage) Append(name, log string) {
	l := LogBody{
		Text: log,
		Name: name,
	}
	s.p.append(fmt.Sprintf("[%s]", name), log)
	s.c.append(name, l)
	for _, c := range s.registee {
		c <- l
	}
}

func (s *Storage) GetCached(name string) []LogBody {
	return s.c.getCached(name)
}

func (s *Storage) GetNames() []string {
	return s.c.getNames()
}

func (s *Storage) Regist(r chan<- LogBody) {
	s.registee = append(s.registee, r)
}
