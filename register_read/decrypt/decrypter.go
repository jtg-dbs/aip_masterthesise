package decrypter

import (
	"log"

	"github.com/billgraziano/dpapi"
	"golang.org/x/sys/windows"
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

// IpSecDll will unprotect a []byte with the ipsecproc.dll of windows OS
func IpSecProcDll() {
	log.SetFlags(0)
	log.SetPrefix("decrypter: ")
	_, err := windows.LoadDLL("ipsecproc.dll")
	if err != nil {
		log.Fatal(err)
	}
}
