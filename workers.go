package main

import (
  "fmt"
)

func f(left chan<- int, right <-chan int) {
    left <- 1 + <-right
}

func main() {
    const n = 100_000

    // first we construct an array of n+1 channels each being a 'chan int'
    var channels [n+1]chan int
    for i := range channels {
        channels[i] = make(chan int)
    }

    // now we wire n goroutines in a chain
    for i := 0; i < n; i++ {
        go f(channels[i], channels[i+1])
    }

    // insert a value into the right-hand end
    go func(c chan<- int) { c <- 1 }(channels[n])

    // pick up the value emerging from the left-hand end
    fmt.Println(<-channels[0])
}
