package main

import (
    "fmt"
    "time"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	a := make(chan string)
	go TickBoom(a)
	for b:= <-a {
		fmt.Println(b)
	}
	wg.Wait()
}

func TickBoom(out chan string, wg *sync.WaitGroup){
	defer wg.Done()
    ticker := time.NewTicker(time.Millisecond * 250)
    boom := time.After(time.Millisecond * 1000)
    for {
        select {
        case <-ticker.C:
            //fmt.Println("tick")
			out <- "tick"
        case <-boom:
            //fmt.Println("boom!")
			out <- "boom!"
            return
		//default:
			//out <- "can't wait."
        }
    }
}