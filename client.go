package main

import (
	"github.com/zhouhui8915/go-socket.io-client"
	"log"
	"time"
)

func client(uri, event string, i int, ch chan struct{}) {

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string), // @todo
	}
	log.Println("new connection")
	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On(event, func(msg string) {
		log.Printf("event: %s message:%v\n", event, msg)
	})
	client.On("disconnection", func() {
		log.Printf("on disconnect #%d \n", i)
		ch <- struct{}{}
	})

	go func() {
		for {
			log.Println("ping...  #", i)
			client.Emit("p", "ping")
			time.Sleep(time.Second * time.Duration(10))
		}
	}()

	time.Sleep(time.Second * time.Duration(waitSecond))
	log.Println("close... #", i)
	ch <- struct{}{}
}
