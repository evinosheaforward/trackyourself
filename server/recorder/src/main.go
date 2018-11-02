package main

import (
  "io/ioutil"
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	"os"
  "time"

	"record"
)

func main() {
	log.Println("***************** START *****************")
	time.Sleep(5*time.Second)
	conn := record.NewDBConn()
	record.SetupDB(conn)
	log.Println("************** DB SETUP ******************")
  http.HandleFunc("/data",
func (rw http.ResponseWriter, req *http.Request) {
	jsonBytes, _ := ioutil.ReadAll(req.Body)
	bodyString := string(jsonBytes)
	log.Println(bodyString)
	conn.Insert(bodyString)
	log.Println("Inserted")
	})
  if err := http.ListenAndServe(os.Getenv("LISTEN_PORT"), nil); err != nil {
    panic(err)
  }
	log.Println("**************** END ********************")
}
