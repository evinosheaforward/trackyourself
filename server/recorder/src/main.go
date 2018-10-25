package main

import (
  "io/ioutil"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func record(rw http.ResponseWriter, req *http.Request) {
		// var anyJson map[string]interface{}
		// decoder := json.NewDecoder(req.Body)
    // err := decoder.Decode(&anyJson)
    // if err != nil {
    //     panic(err)
    // }
    // log.Println(anyJson["test"])
		// textJson, _ := json.Marshal(anyJson)
		textJson, _ := ioutil.ReadAll(req.Body)
		log.Println(textJson)
		rightNow := time.Now().String()
		outName := os.Getenv("DATA_DIR") + rightNow + ".json"
		log.Println(outName)
		_ = ioutil.WriteFile(outName, textJson, 0644)
}

func helloWorld(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello World!")
}

func main() {
	// http.HandleFunc("/data", helloWorld)
  // if err := http.ListenAndServe(":9299", nil); err != nil {
  //   panic(err)
  // }
  http.HandleFunc("/data", record)
  if err := http.ListenAndServe(":9299", nil); err != nil {
    panic(err)
  }
}
