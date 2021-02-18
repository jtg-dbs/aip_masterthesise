package decrypter

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

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
func IpSecProcDll(key []byte) []byte {
	log.SetFlags(0)
	log.SetPrefix("decrypter: ")
	// Create temp files and directories
	dir, err := ioutil.TempDir("./decrypt", "temp")
	check(err, "create Directory: "+dir)
	file, err := ioutil.TempFile(dir, "sk_unencrypted")
	check(err, "creating Temp File")
	_, err = file.Write(key)
	check(err, "create sk-encrypted file")
	// execute decrypter.exe
	exec_decrypt := exec.Command("powershell.exe", `.\decrypt\decrypter.exe`, `.\`+file.Name()+" "+`.\`+dir+`\sk-rsa-decrypted.dat`)
	err = exec_decrypt.Run()
	check(err, "execute decrypter.exe")
	// Read decrypted sk file
	sk_decrypted, err := ioutil.ReadFile(`.\` + dir + `\sk-rsa-decrypted.dat`)
	check(err, "read decrytped sk file")
	file.Close()
	//delte temp
	remove_dir_content(dir)
	return sk_decrypted
}

func check(e error, s string) {
	if e != nil {
		log.Println(s)
		log.Fatal(e)
	}
}

func remove_dir_content(dir string) {
	files, err := ioutil.ReadDir(dir)
	check(err, "delte temp files")
	for _, d := range files {
		os.RemoveAll(path.Join([]string{dir, d.Name()}...))
	}
	err = os.Remove(dir)
	check(err, "delte temp directory")
}
