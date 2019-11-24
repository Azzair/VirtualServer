package main

import "fmt"

import "net"

import "log"

import "io"

func main() {
	fmt.Println("Запускаю додаток.....")

	li, err := net.Listen("tcp", ":10101")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		log.Println("New connection: ", conn.RemoteAddr())
		if err != nil {
			log.Fatalln(err)
		}
		io.WriteString(conn, "\n Hello from TCP\n")
		fmt.Fprintln(conn, "fmt.Hello")
		log.Println("Відповідь надіслано")
		conn.Close()
	}
}
