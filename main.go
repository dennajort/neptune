package main

import (
	"log"
	"os"

	"github.com/gossel-j/neptune/src/bencode"
)

func printInt(i int64) {
	println(i)
}

func main() {
	data, err := bencode.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
	err = bencode.Marshal(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
