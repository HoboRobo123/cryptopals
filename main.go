package main

import (
	"fmt"
	HexToBase64 "hexToBase64"
)

func main() {
	result := HexToBase64.Convert("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println(result)
}
