package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
const address = "localhost:8080"

func main() {
	go runServer()

	timeOut := flag.Duration("timeout", 10*time.Second, "Time out flag")
	flag.Parse()
	time.Sleep(time.Second)
	conn, err := net.DialTimeout("tcp", address, *timeOut)
	if err != nil {
		log.Fatalln(err)
	}
	go writeToSocket(conn)

	for {
		text, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed")
			break
		}
		fmt.Println("Get message from the socket: ", text)
	}

}
func writeToSocket(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("Got ctrl+D signal.Closing connection")
			conn.Close()
			return
		}
		_, err = conn.Write(text)
		if err != nil {
			conn.Close()
			fmt.Println("Connection closed due to error:", err)
			return
		}
	}
}
func runServer() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := listener.Accept()
	if err != nil{
		return
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	for {
		text, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}
		fmt.Println("Server get message: ", text)

		conn.Write([]byte("New server answer: " + text + "\n"))
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}
	}
}