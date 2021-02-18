package output

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
)

type export struct {
	Sk string
	Mk string
}

func Export(sk []byte, mk []byte) {
	log.SetFlags(0)
	log.SetPrefix("export: ")

	output := export{
		Sk: hex.EncodeToString(sk),
		Mk: hex.EncodeToString(mk)}
	output_json, err := json.Marshal(output)
	check(err, "Build JSON Object")
	err_write := ioutil.WriteFile("sk_mk.json", output_json, 0777)
	check(err_write, "Writing of the output")
}

func check(e error, s string) {
	if e != nil {
		log.Println(s)
		log.Fatal(e)
	}
}
