package main

import (
        "fmt"
        "sync"
        "time"
)

var wg sync.WaitGroup
var mu sync.Mutex
var sum = 0
var times = 0
var pass_message = 11

func process(n int, c chan int) {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            times = times + 1
            sum = sum + pass_message
            mu.Unlock()
            time.Sleep(100 * time.Millisecond)
            fmt.Println("From ", n+1)
            c <- pass_message // send sum to c	
            pass_message += 100000
        }()
}


var pass_message3 = 0
var N = 10
func main() {
    for j := 0; j < N; j++ {
        
        c := make(chan int)
        
        go process(j, c)
        
        pass_message2 := <-c // receive from c
        
        pass_message3 = pass_message2
        
	    fmt.Println(pass_message2)
	    
        wg.Wait()

    }
    fmt.Println("Message was :", pass_message3-((N-1)*100000))
    fmt.Println("Number of processes :", (pass_message3 - (pass_message3-((N-1)*100000)))/100000 + 1)
    fmt.Println("Message using Locks(as sum of pass_messages from every iteration):", sum, ".  Temes of work: ", times )
}