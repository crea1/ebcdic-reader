package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
)

func main() {

	fmt.Println("Hello :)")
	file, err := os.Open("TEST.EBC")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		} else {

			fmt.Println(b)
		}

	}
}
