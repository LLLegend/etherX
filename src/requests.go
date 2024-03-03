package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func DebugTransaction(client *http.Client, hash string) (string, error) {
	json := `{"id": 1, "method": "debug_traceTransaction", "params": ["` + hash + `", {"tracer": "callTracer", "tracerConfig": {"onlyTopCall": false, "withLog":true}}]}`
	log.Println("Request with:", json)
	jsonByte := []byte(json)
	req, _ := http.NewRequest("POST", "127.0.0.1:8545", bytes.NewBuffer(jsonByte))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("Request info: " + resp.Status + " " + string(body))
	return string(body), err
}
