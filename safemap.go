package main

import (
	"fmt"
	"sync"
	"time"

)

/*

golang safemap implemenatation
	normal map is not rottine safe(thread safe ) if multiple threads are using same map
	will lead to race condtion

	so we have to implement a mutex lock to avoid this

*/

// init a map
type SafeMap[K comparable, V any] struct {
	lock sync.RWMutex
	Map  map[K]V
}

// create a function that will instantiate a new map

func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		Map: make(map[K]V),
	}
}

// set value function
func (M *SafeMap[K, V]) Set(Key K, value V) {
	M.lock.Lock()
	defer M.lock.Unlock()
	M.Map[Key] = value
}

// get value
func (M *SafeMap[K, V]) Get(Key K) (V, error) {
	M.lock.RLock()
	defer M.lock.RUnlock()
	val, ok := M.Map[Key]
	if ok {
		return val, nil
	}
	return val, fmt.Errorf("invalid Key")
}

func main() {
	new_map := New[int, string]()

	for i := 0; i < 10; i++ {
		go func(i int) {
			new_map.Set(i, "hello")
		}(i)
	}
	time.Sleep(time.Second * 4)
	fmt.Println(new_map)
}
