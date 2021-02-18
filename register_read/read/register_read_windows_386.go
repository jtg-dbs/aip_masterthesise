package read

import (
	error_log "registryread/error"

	"golang.org/x/sys/windows/registry"
)

// This package contains the functions to read MK und SK from the registry
// Furhtermore it will extract the hexbits of the registry values

// OpenMkSk extracts the mk and sk value for a specific user and returns an error if needed
func OpenMkSK() ([]byte, []byte) {
	// Open registry Key
	msippKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\Local Settings\Software\Microsoft\MSIPC`, registry.QUERY_VALUE)
	if err != nil {
		error_log.Check(err, "cant open registry", "open_mk_sk")
		msippKey.Close()
	}

	// Read MSIPP-MK Value from Registry Key
	msippMkBin, _, err := msippKey.GetBinaryValue("MSIPP-MK")
	error_log.Check(err, "cant read MK Value", "open_mk_sk")

	//read MSIPP-SK Value from Registry Key
	msippSkBin, _, err := msippKey.GetBinaryValue("MSIPP-SK")
	error_log.Check(err, "cant read SK Value", "open_mk_sk")
	return msippSkBin, msippMkBin
}
