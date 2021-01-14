package main

import (
	"fmt"
	"log"
	"registerread/read"
)

func main() {
	msippSkBin, msippMkBin := read.OpenMkSK()
	log.SetFlags(0)
	log.SetPrefix("main: ")

	log.Print("Sucesfully read MK and SK from registry")
	fmt.Printf("SK value is: \n %X \n", msippSkBin)
	fmt.Printf("MK value is: \n %X \n", msippMkBin)
}
