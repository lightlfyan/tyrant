package data

import "sync"

type TaskProgressType struct {
	taskProgress map[int32]int32
	lock         *sync.RWMutex
}

func (t *TaskProgressType) Init() {
	t.taskProgress = make(map[int32]int32, MAX_WORKER)
	t.lock = &sync.RWMutex{}

	for i := 0; i < MAX_WORKER; i++ {
		t.Set(int32(i), 0)
	}
}

func (t *TaskProgressType) Get(id int32) (int32, bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	v, ok := t.taskProgress[id]
	return v, ok
}

func (t *TaskProgressType) Set(k, v int32) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.taskProgress[k] = v
}
