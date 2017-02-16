package main

import (
	"log"
	"os"

	"gitlab.com/dennajort/neptune/bencode"
)

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
