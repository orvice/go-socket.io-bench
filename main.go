package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	addr       = "localhost"
	clientNum  = 100
	waitSecond = 100
	customOn   = "message"
)

func initFlag() {
	flag.IntVar(&clientNum, "n", 10, "num of client")
	flag.IntVar(&waitSecond, "s", 1024, "second for wait client")
	flag.StringVar(&addr, "addr", "localhost", "socket.io addr")
	flag.StringVar(&customOn, "on", "message", "custom listen event")
	flag.Parse()
}

func main() {
	initFlag()
	fmt.Printf("start test addr: %s  client num: %d  wait second: %d  ", addr, clientNum, waitSecond)

	var wg sync.WaitGroup

	for i := 0; i < clientNum; i++ {
		wg.Add(1)
		fmt.Println("start #", i)
		go func(i int) {
			ch := make(chan struct{})
			go client(addr, customOn, ch)
			<-ch
			log.Println("done #", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	log.Println("done....exit")
	os.Exit(0)
}
