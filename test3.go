package main

import (  
    "fmt"
    "sync"
)

func main() {  
    var wg sync.WaitGroup
    done := make(chan struct{})
    workerCount := 2

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go doit(i,done,&wg)
    }
	fmt.Println("all go dispatched, closing done chan")
    close(done)
    wg.Wait()
    fmt.Println("all done!")
}

func doit(workerId int,done <-chan struct{},wg *sync.WaitGroup) {  
    fmt.Printf("[%v] is running\n",workerId)
    defer wg.Done()
	fmt.Println("getting from done...")
    <- done
    fmt.Printf("[%v] is done\n",workerId)
}