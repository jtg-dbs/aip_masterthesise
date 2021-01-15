package main

import (
	"fmt"
	"log"
	decrypter "registryread/decrypt"
	"registryread/read"
)

func main() {
	msippSkBin, msippMkBin := read.OpenMkSK()
	log.SetFlags(0)
	log.SetPrefix("main: ")

	log.Print("Sucesfully read MK and SK from registry")
	MkBinDec := decrypter.Decrypt(msippMkBin)
	SkBinDec := decrypter.Decrypt(msippSkBin)

	log.Print("Sucesfully decrpyted with DPAPI")
	fmt.Printf("%X", MkBinDec)
	fmt.Printf("%X", SkBinDec)
}
