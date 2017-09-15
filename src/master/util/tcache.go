package util

import "sync"

type TCache struct {
	Dict map[int64]interface{}
	Lock *sync.RWMutex
}

func (t *TCache) MaxKey() int64 {
	var i int64 = 0

	t.Lock.RLock()
	defer t.Lock.RUnlock()

	for k, _ := range t.Dict {
		if k > i {
			i = k
		}
	}
	return i
}

func (t *TCache) Set(k int64, r interface{}) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.Dict[k] = r
}

func (t *TCache) Get(k int64) (interface{}, bool) {
	t.Lock.Lock()
	defer t.Lock.Unlock()

	r, ok := t.Dict[k]
	return r, ok
}

func (t *TCache) UnsafeDel(k int64) {
	delete(t.Dict, k)
}

func (t *TCache) Del(k int64) {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	delete(t.Dict, k)
}

func (t *TCache) UnsafeSet(k int64, r interface{}) {
	t.Dict[k] = r
}

func (t *TCache) ForeachRead(F func(int64, interface{})) {
	t.Lock.RLock()
	defer t.Lock.RUnlock()

	for k, v := range t.Dict {
		F(k, v)
	}
}

func (t *TCache) ForeachWrite(F func(int64, interface{})) {
	t.Lock.RLock()
	defer t.Lock.RUnlock()

	for k, v := range t.Dict {
		F(k, v)
	}
}

func (t *TCache) Reset() {
	t.Lock.Lock()
	defer t.Lock.Unlock()
	t.Dict = make(map[int64]interface{}, 0)
}
