package main

import (
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

var (
	addr        = flag.String("addr", "ws://127.0.0.1:6060/", "websocket addr")
	connections = flag.Int("num", 1000, "number of websocket connections")

	total   int64
	success int64
	fail    int64
	alive   int64
)

func main() {
	flag.Usage = func() {
		io.WriteString(os.Stderr, `Websockets client generator
Example usage: ./bench -addr=ws://127.0.0.1:6060 -num=1000
`)
		flag.PrintDefaults()
	}
	flag.Parse()
	target := *addr
	log.Printf("Connecting to %s", target)
	for i := 0; i < *connections; i++ {
		go connect()
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Hour)
}

func connect() {
	n := atomic.AddInt64(&total, 1)
	if n%100 == 0 {
		log.Printf("connections %d ", n)
	}
	c, _, err := websocket.DefaultDialer.Dial(*addr, nil)
	if err != nil {
		n := atomic.AddInt64(&fail, 1)
		if n%100 == 0 {
			log.Printf("fail %d %s", n, err)
		}
		return
	}
	n = atomic.AddInt64(&success, 1)
	a := atomic.AddInt64(&alive, 1)
	defer func() {
		a := atomic.AddInt64(&alive, -1)
		if a%100 == 0 {
			log.Printf("alive %d ", n)
		}
	}()
	sIndex := strconv.Itoa(rand.Intn(10))
	err = c.WriteMessage(websocket.TextMessage, []byte("{\"listen\":true,\"event\":\"v-"+sIndex+"\"}"))
	if err != nil {
		log.Printf("msg error %s", err)
	}
	if n%100 == 0 {
		log.Printf("success %d ", n)
	}
	if a%100 == 0 {
		log.Printf("alive %d ", a)
	}
	for {
		time.Sleep(time.Minute)
		err = c.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second*5))
		if err != nil {
			return
		}
	}
}
