package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Articles/Torii/internal/client"
)

func main() {
	c, err := client.New(":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Connected to server. Type messages:")

	for {
		fmt.Print(">> ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("read error:", err)
			return
		}

		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("bye 👋")
			return
		}

		resp, err := c.Send(text)
		if err != nil {
			log.Println("send error:", err)
			return
		}

		fmt.Println("echo:", strings.TrimSpace(resp))
	}
}

