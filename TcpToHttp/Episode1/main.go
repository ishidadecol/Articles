// TcpToHttp/Episode1/main.go
package main

import (
	"fmt"
	"io"
	"os" // package que vamos utilizar para lermos o arquivo
)

func main() {

	filepath := "../message.txt" //path relativo para o txt
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	buffer := make([]byte, 8)

	for {
		byteCount, err := file.Read(buffer)
		if byteCount > 0 {
			fmt.Printf("READ: %s\n", string(buffer[:byteCount]))
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
	}
}

