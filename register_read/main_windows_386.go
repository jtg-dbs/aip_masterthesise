package main

import (
	decrypter "registryread/decrypt"
	error_log "registryread/error"
	"registryread/output"
	"registryread/read"
)

func main() {
	msippSkBin, msippMkBin := read.OpenMkSK()
	error_log.Log("main", "Sucesfully read MK and SK from registry")

	MkBinDec := decrypter.Decrypt(msippMkBin)
	SkBinDec := decrypter.Decrypt(msippSkBin)
	error_log.Log("main", "Sucesfully decrpyted with DPAPI")

	sk := decrypter.IpSecProcDll(SkBinDec)
	error_log.Log("main", "sucesfully decrypt sk with IPCSECPROC")

	output.Export(sk, MkBinDec)
	error_log.Log("main", "Sucesfully exported sk and mk as JSON")
}
