package main

import (
  "io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type test_struct struct {
    Test string
}

func record(rw http.ResponseWriter, req *http.Request) {
		var anyJson map[string]interface{}
		decoder := json.NewDecoder(req.Body)
    err := decoder.Decode(&anyJson)
    if err != nil {
        panic(err)
    }
    log.Println(anyJson["test"])
		textJson, _ := json.Marshal(anyJson)
		rightNow := time.Now().String()
		outName := os.Getenv("DATA_DIR") + rightNow + ".json"
		log.Println(outName)
		err = ioutil.WriteFile(outName, textJson, 0644)
}

func main() {
  http.HandleFunc("/data", record)
  if err := http.ListenAndServe(":9299", nil); err != nil {
    panic(err)
  }
}
