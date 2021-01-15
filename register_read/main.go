package main

import (
	"log"
	decrypter "registryread/decrypt"
	"registryread/read"
)

func main() {
	msippSkBin, msippMkBin := read.OpenMkSK()
	log.SetFlags(0)
	log.SetPrefix("main: ")

	log.Print("Sucesfully read MK and SK from registry")
	decrypter.Decrypt(msippMkBin)
	decrypter.Decrypt(msippSkBin)
}
