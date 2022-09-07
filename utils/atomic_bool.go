package utils

import "sync"

type AtomicBool struct {
	mutex sync.Mutex
	value bool
}

func (atomicBool *AtomicBool) CompareAndSet(expected bool, new bool) bool {
	atomicBool.mutex.Lock()
	defer atomicBool.mutex.Unlock()

	if atomicBool.value == expected {
		atomicBool.value = new
		return true
	} else {
		return false
	}

}
