package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
	"flag"
)

func main() {
	var ebcdicFile string = ""

	var use277 = flag.Bool("277", false, "Wether to use EBCDIC 277 code page")
	flag.Parse()

	if len(flag.Args()) == 1 {
		ebcdicFile = flag.Args()[0]
	} else {
		fmt.Println("Error: No filename specified")
		fmt.Println("\nUsage:\n    " + os.Args[0] + "[-277] <EBCDIC file>")
		os.Exit(1)
	}

	file, err := os.Open(ebcdicFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var result []rune
	byteCounter := 0

	var ebcdic_to_ascii_map map[byte]byte
	if *use277 {
		ebcdic_to_ascii_map = ebcdic_to_ascii_277
	} else {
		ebcdic_to_ascii_map = ebcdic_to_ascii
	}

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

			if val, ok := ebcdic_to_ascii_map[b]; ok {
				result = append(result, rune(val))
			}

		}

	}

	//fmt.Printf("Read bytes : %d Output bytes : %d\n", byteCounter, len(result))
	fmt.Println(string(result[:]))
}
