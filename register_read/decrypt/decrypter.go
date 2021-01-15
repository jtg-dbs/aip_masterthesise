package decrypter

import (
	"log"

	"github.com/billgraziano/dpapi"
)

// Decrypt will unprotect []byte with the DPAPI in context of the current user
func Decrypt(key []byte) []byte {
	log.SetFlags(0)
	log.SetPrefix("decrypter: ")
	decryptedbin, err := dpapi.DecryptBytes(key)
	if err != nil {
		log.Fatal(err)
	}
	return decryptedbin
}