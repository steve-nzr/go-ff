package main

import (
	"fmt"
	"net"
	"time"
)

func load() {
	conn, err := net.Dial("tcp", "localhost:5400")
	if err != nil {
		return
	}

	conn.Write([]byte("kdfjgsdfjgsdfjgskdfjgksdjfgkjsdkfjgskdfjgkj"))
	conn.Write([]byte("kdfjgsdfjgsdfjgskdfjgksdjfgkjsdkfjgskdfjgkj"))
	conn.Write([]byte("kdfjgsdfjgsdfjgskdfjgksdjfgkjsdkfjgskdfjgkj"))
	conn.Write([]byte("kdfjgsdfjgsdfjgskdfjgksdjfgkjsdkfjgskdfjgkj"))
	conn.Write([]byte("kdfjgsdfjgsdfjgskdfjgksdjfgkjsdkfjgskdfjgkj"))
	conn.Write([]byte("kdfjgsdfjgsdfjgskdfjgksdjfgkjsdkfjgskdfjgkj"))
	conn.Write([]byte("kdfjgsdfjgsdfjgskdfjgksdjfgkjsdkfjgskdfjgkj"))
	conn.Write([]byte("kdfjgsdfjgsdfjgskdfjgksdjfgkjsdkfjgskdfjgkj"))
}

func main() {
	for i := 0; i < 100000; i++ {
		go load()
	}

	fmt.Println("Done")
	time.Sleep(20 * time.Second)
}
