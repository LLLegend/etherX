package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func doRequest(client *http.Client, request *http.Request) (map[string]interface{}, error) {
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Request Error:", err)
		return make(map[string]interface{}), err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read Json Error: ", err)
		return make(map[string]interface{}), err
	}

	var bodyJsonMap map[string]interface{}
	d := json.NewDecoder(strings.NewReader(string(body)))
	d.UseNumber()
	err = d.Decode(&bodyJsonMap)
	if err != nil {
		fmt.Println("Decode Json Error :", err)
		return make(map[string]interface{}), err
	}

	//err = json.Unmarshal(body, &bodyJsonMap)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
	return bodyJsonMap, err
}
