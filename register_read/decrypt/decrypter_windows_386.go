package decrypter

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/billgraziano/dpapi"
)

// Decrypt will unprotect []byte with the DPAPI in context of the current user
func Decrypt(key []byte) []byte {
	log.SetFlags(0)
	log.SetPrefix("decrypter: ")
	decryptedbin, err := dpapi.DecryptBytes(key)
	check(err, "Unprotect String with DPAPI")
	return decryptedbin
}

// IpSecProcDll will unprotect a []byte with the ipsecproc.dll of windows OS
func IpSecProcDll(key []byte) {
	log.SetFlags(0)
	log.SetPrefix("decrypter: ")
	dir, err := ioutil.TempDir("./decrypt", "temp")
	base_dir := filepath.Base(dir)
	check(err, "create Directory: "+dir)
	file, err := ioutil.TempFile(dir, "sk_unencrypted")
	base_file := filepath.Base(file.Name())
	check(err, "creating Temp File")
	_, err = file.Write(key)
	check(err, "create sk-encrypted file")
	exec_decrypt := exec.Command("powershell.exe", `.\decrypt\decrypter.exe`, `.\`+file.Name()+" "+`.\`+dir+`\sk-rsa-decrypted.dat`)
	err = exec_decrypt.Run()
	check(err, "execute decrypter.exe")
}

func check(e error, s string) {
	if e != nil {
		fmt.Printf(s + "\n")
		log.Fatal(e)
	}
}
