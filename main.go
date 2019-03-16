package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/danward79/BOM/bom"
)

func main() {

	var o bom.StationData

	file, err := ioutil.ReadFile("testdata/fawkner.json")
	doError(err)

	//fmt.Println("file:", string(file))

	err = json.Unmarshal(file, &o)
	doError(err)

	fmt.Println(o.Observations.Header[0].RefreshMessage)
}

//doError lazy error handling
func doError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
