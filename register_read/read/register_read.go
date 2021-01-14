package read

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

// This package contains the functions to read MK und SK from the registry
// Furhtermore it will extract the hexbits of the registry values

// OpenMkSk extracts the mk and sk value for a specific user and returns an error if needed
func OpenMkSK() ([]byte, []byte) {
	log.SetFlags(0)
	log.SetPrefix("open_mk_sk: ")

	// Open registry Key
	msippKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\Local Settings\Software\Microsoft\MSIPC`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
		msippKey.Close()
	}

	// Read MSIPP-MK Value from Registry Key
	msippMkBin, _, err := msippKey.GetBinaryValue("MSIPP-MK")
	if err != nil {
		log.Fatal(err)
	}

	//read MSIPP-SK Value from Registry Key
	msippSkBin, _, err := msippKey.GetBinaryValue("MSIPP-SK")
	if err != nil {
		log.Fatal(err)
	}
	return msippSkBin, msippMkBin

}
