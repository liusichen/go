package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type serverList struct {
	ID             string
	SN             string
	IDC            string
	City           string       `json:"city"`
	Status         string       `json:"status"`
	Manufacturer   string       `json:"manufacturer"`
	Model          string       `json:model`
	Service        []string     `json:service`
	ServiceAlias   []string     `json:"service_alias"`
	Isp            string       `json:"isp"`
	PrivateAddress []string     `json:"private_address"`
	PublicAddress  []string     `json:"public_address"`
	DracAddress    string       `json:"drac_address"`
	VirtualAddress []string     `json:"virtual_address"`
	NewSer         []newService `json:"new_service"`
}

type newService struct {
	name string
	code string
}

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var snip []struct {
		SN            string
		PublicAddress []string `json:"public_address"`
	}
	err = json.Unmarshal(data, &snip)
	for _, sn := range snip {
		fmt.Println(sn.SN, sn.PublicAddress)
	}
}
