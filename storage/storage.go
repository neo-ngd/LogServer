package storage

import (
	"time"
)

type LogBody struct {
	Text string
	Name string
	Key  int64
}

type Storage struct {
	p *persist
	c *logcache
}

func NewStorage(path, name string, age, split time.Duration) *Storage {
	return &Storage{
		p: newPersist(path, name, age, split),
		c: newLogCache(50),
	}
}
func (s *Storage) Append(name, log string) {
	l := LogBody{
		Text: log,
		Name: name,
	}
	s.p.append(name, log)
	s.c.append(name, l)
}

func (s *Storage) GetCached(name string) []LogBody {
	return s.c.getCached(name)
}

func (s *Storage) GetNames() []string {
	return s.c.getNames()
}
