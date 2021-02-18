package decrypter

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	error_log "registryread/error"

	"github.com/billgraziano/dpapi"
)

// Decrypt will unprotect []byte with the DPAPI in context of the current user
func Decrypt(key []byte) []byte {
	decryptedbin, err := dpapi.DecryptBytes(key)
	error_log.Check(err, "Unprotect String with DPAPI", "decrypter")
	return decryptedbin
}

// IpSecProcDll will unprotect a []byte with the ipsecproc.dll of windows OS
func IpSecProcDll(key []byte) []byte {
	// Create temp files and directories
	dir, err := ioutil.TempDir("./decrypt", "temp")
	error_log.Check(err, "create Directory: "+dir, "decrypter")
	file, err := ioutil.TempFile(dir, "sk_unencrypted")
	error_log.Check(err, "creating Temp File", "decrypter")
	_, err = file.Write(key)
	error_log.Check(err, "create sk-encrypted file", "decrypter")
	// execute decrypter.exe
	exec_decrypt := exec.Command("powershell.exe", `.\decrypt\decrypter.exe`, `.\`+file.Name()+" "+`.\`+dir+`\sk-rsa-decrypted.dat`)
	err = exec_decrypt.Run()
	error_log.Check(err, "execute decrypter.exe", "decrypter")
	// Read decrypted sk file
	sk_decrypted, err := ioutil.ReadFile(`.\` + dir + `\sk-rsa-decrypted.dat`)
	error_log.Check(err, "read decrytped sk file", "decrypter")
	file.Close()
	//delte temp
	remove_dir_content(dir)
	return sk_decrypted
}

func remove_dir_content(dir string) {
	files, err := ioutil.ReadDir(dir)
	error_log.Check(err, "delte temp files", "decrypter")
	for _, d := range files {
		os.RemoveAll(path.Join([]string{dir, d.Name()}...))
	}
	err = os.Remove(dir)
	error_log.Check(err, "delte temp directory", "decrypter")
}
