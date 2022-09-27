package queue

import "fmt"

type Queue struct {
	ch chan int
}

func newQueue() *Queue {
	return &Queue{
		ch: make(chan int),
	}
}

func (q *Queue) Produce(val int) {
	q.ch <- val
}

func (q *Queue) Consume() (int, bool) {
	val, ok := <-q.ch
	return val, ok
}

func (q *Queue) Close() {
	close(q.ch)
}

func produce(q *Queue) {
	for i := 0; i < 10; i++ {
		fmt.Println("produce", i)
		q.ch <- i
	}
	q.Close()
}

func consume(q *Queue) {
	for val := range q.ch {
		fmt.Println("consume", val)
	}
}

func Run() {
	q := newQueue()
	go produce(q)
	consume(q)
	fmt.Println("done")
}
