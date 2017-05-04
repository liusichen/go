package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var keyValue []struct {
	SN             string
	PrivateAddress []string `json:"private_address"`
	PublicAddress  []string `json:"public_address"`
}

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	for _, url := range os.Args[1:] {
		resq, err := client.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		data, _ := ioutil.ReadAll(resq.Body)
		defer resq.Body.Close()
		err = json.Unmarshal(data, &keyValue)
		for _, value := range keyValue {

			fmt.Println(value.SN, value.PrivateAddress, value.PublicAddress)
		}
	}
}
