package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	// Define the string to encode
	s := "Hello, World!"

	// Create a byte buffer to hold the encoded data
	buf := make([]byte, len(s)*2)

	// Encode the string to UTF-16
	for i, r := range s {
		binary.LittleEndian.PutUint16(buf[i*2:], uint16(r))
	}

	// Write the encoded data to a file
	f, err := os.Create("utf16_data.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("UTF-16 data written to utf16_data.bin")
}
