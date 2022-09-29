package cpu

import (
	"fmt"
	"sync"
)

var Ctx = sync.WaitGroup{}

func Run(n int) {
	var i uint64
	for ; ; i++ {
		if i == 0 {
			fmt.Println("run again", n)
		}
	}
}
