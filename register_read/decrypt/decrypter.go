package decrypter

import "fmt"

// Decrypt will unprotect []byte with the DPAPI in context of the current user
func Decrypt(key []byte) {
	fmt.Printf("%X", key)
}
