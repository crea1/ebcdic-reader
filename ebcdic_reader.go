package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
)

func main() {
	file, err := os.Open("TEST.EBC")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var result []byte
	byteCounter := 0

	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		} else {
			byteCounter++
			// For some reason I only get CR and not CRLF from EBCDIF file. Add LF for printing
			if b == 0x0d {
				result = append(result, 0x0d, 0x0a)
			}
			if val, ok := ebcdic_to_ascii[b]; ok {
				result = append(result, val)
			}


		}

	}

	//fmt.Printf("Read bytes : %d Output bytes : %d\n", byteCounter, len(result))
	fmt.Println(string(result[:]))
}
