package output

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	error_log "registryread/error"
)

type export struct {
	Sk string
	Mk string
}

func Export(sk []byte, mk []byte) {
	output := export{
		Sk: hex.EncodeToString(sk),
		Mk: hex.EncodeToString(mk)}
	output_json, err := json.Marshal(output)
	error_log.Check(err, "Build JSON Object", "export")
	err_write := ioutil.WriteFile("sk_mk.json", output_json, 0777)
	error_log.Check(err_write, "Writing of the output", "export")
}
