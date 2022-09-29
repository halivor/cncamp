package queue

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/halivor/cncamp/module2/internal/queue"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "queue",
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}
}

func Run() {
	q := queue.New()
	for i := 0; i < 3; i++ {
		go Consume(i+1, q)
	}
	Produce(q)
}

func Produce(q *queue.Queue) {
	tk := time.NewTicker(time.Second)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			for i := rand.Intn(9); i > 0; i-- {
				q.Produce(strconv.Itoa(rand.Int()))
			}
			q.ActiveConsumer()
		}
	}
}

func Consume(no int, q *queue.Queue) {
	for {
		fmt.Println(no, "consume", q.Consume())
	}
}
