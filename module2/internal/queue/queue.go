package queue

import (
	"fmt"
	"sync"
)

type Queue struct {
	lock  sync.Mutex
	cond  *sync.Cond
	queue []string
}

func New() *Queue {
	Q := &Queue{
		queue: make([]string, 0, 128),
	}
	Q.cond = sync.NewCond(&Q.lock)
	return Q
}

func (q *Queue) Produce(val string) {
	q.lock.Lock()
	defer q.lock.Unlock()
	fmt.Println("produce", val)
	q.queue = append(q.queue, val)
}

func (q *Queue) ActiveConsumer() {
	q.cond.Signal()
}

func (q *Queue) Consume() string {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) == 0 {
		q.cond.Wait()
	}
	if len(q.queue) == 0 {
		return ""
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}
