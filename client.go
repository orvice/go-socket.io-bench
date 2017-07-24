package main

import (
	"bufio"
	"github.com/zhouhui8915/go-socket.io-client"
	"log"
	"os"
)

func client(uri, event string) {

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string), // @todo
	}

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
		log.Printf("on disconnect\n")
	})

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		client.Emit("message", command)
		log.Printf("send message:%v\n", command)
	}
}
