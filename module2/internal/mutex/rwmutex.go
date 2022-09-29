package mutex

import (
	"fmt"
	"sync"
)

var (
	lock = sync.RWMutex{}
	wg   = sync.WaitGroup{}
)

func Run() {
	wg.Add(2)
	go rLock()
	go wLock()
	wg.Wait()
}

func rLock() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rlock", i)
	}
}

func wLock() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		func() {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("lock", i)
		}()
	}
}
