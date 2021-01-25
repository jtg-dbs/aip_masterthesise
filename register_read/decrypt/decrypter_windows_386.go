package decrypter

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

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

// IpSecProcDll will unprotect a []byte with the ipsecproc.dll of windows OS
func IpSecProcDll(key []byte) {
	log.SetFlags(0)
	log.SetPrefix("decrypter: ")

	// Depend the Headers of the sk-unproteted
	skBinProtected := key[4:]

	// Load DLL and make iit usable
	ipcSecProcDll := syscall.NewLazyDLL(`C:\Users\eve.aip\ipcsecproc.dll`)
	err := ipcSecProcDll.Load()
	if err != nil {
		fmt.Print("Load funktion")
		log.Fatal(err)
	}
	proc := ipcSecProcDll.NewProc("IpcSPEncrypt")

	skPointer := unsafe.Alignof(skBinProtected)

	ret, _, err := proc.Call(uintptr(skPointer))
	if err != nil {
		fmt.Printf("%X", ret)
		fmt.Print("fehlermeldung")
		log.Fatal(err)
	}

	fmt.Print(ret)

	// // Loading de ipcsecproc.dll for deobfuscating the skBinProtected
	// ipcSecProcDll, err := windows.LoadDLL("ipsecproc.dll")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
