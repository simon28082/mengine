package pool

import "sync"

type SyncMapPool struct {
	pool sync.Map
}

func NewSyncMapPool() *SyncMapPool {
	return &SyncMapPool{}
}

func (s *SyncMapPool) Get(key any) any {
	v, ok := s.pool.Load(key)
	if ok {
		return v
	}
	return nil
}

func (s *SyncMapPool) Add(key, val any) bool {
	_, ok := s.pool.LoadOrStore(key, val)
	return !ok
}

func (s *SyncMapPool) Put(key, val any) {
	s.pool.Store(key, val)
}

func (s *SyncMapPool) Exists(key any) bool {
	_, ok := s.pool.Load(key)
	return ok
}

func (s *SyncMapPool) Remove(key any) {
	s.pool.Delete(key)
}

func (s *SyncMapPool) Length() (i int) {
	s.pool.Range(func(key, value any) bool {
		i += 1
		return true
	})
	return
}

func (s *SyncMapPool) Clean() {
	s.pool.Range(func(key, value any) bool {
		s.pool.Delete(key)
		return true
	})
}
