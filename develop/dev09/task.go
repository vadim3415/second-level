package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// http://stackoverflow.com/
// Реализовать утилиту wget с возможностью скачивать сайты целиком.

func main() {
	url := os.Args[1]
	resp, err := http.Get(strings.TrimSpace(url))
	if err != nil {
		log.Fatal(err)
	}

	nameFile := strings.Split(url, "/")

	readResp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(fmt.Sprintf("%s.html", nameFile[2]))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(string(readResp))

}
