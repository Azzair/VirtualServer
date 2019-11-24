package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("Запускаю додаток.....")

	li, err := net.Listen("tcp", ":10101")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go hendler(conn)
	}
}

func hendler(conn net.Conn) {
	defer conn.Close()
	err := conn.SetDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("New connection: ", conn.RemoteAddr())
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		ln := scan.Text()
		fmt.Println(ln)
		fmt.Fprintln(conn, "Ohh ....\nHello from SERVER")
	}
	log.Println("Опрацьовано\n\n")
}
