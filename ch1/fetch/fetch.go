package main

import (
	bdutils "bd-utils"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var IPToSNMap = make(map[string]string)
var SNToipMap = make(map[string]map[string][]string)

type wantValue struct {
	SN             string
	PrivateAddress []string `json:"private_address"`
	PublicAddress  []string `json:"public_address"`
}

var keyValue []wantValue

func main() {
	for _, url := range os.Args[1:] {
		urlSlice := strings.Split(url, "/")
		if bdutils.InSlice("boxinfowithidc", urlSlice) == true {
			boxinfoURL(url)
		} else if bdutils.InSlice("?appkey=sync_vidc", urlSlice) == true {
			//TODO:iptoWeigt and weightoIP map
		} else {
			log.Println("the URL cannot be handle.")
		}
	}
}

func boxinfoURL(url string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resq, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	data, _ := ioutil.ReadAll(resq.Body)
	defer resq.Body.Close()
	err = json.Unmarshal(data, &keyValue)
	for _, value := range keyValue {
		SetIPtoSNMap(value)
		SetSNtoIPMap(value)
	}
	for IP, SN := range IPToSNMap {
		fmt.Printf("%s\t%s\n", IP, SN)
	}
	for SN, IPs := range SNToipMap {
		for IPtype, IPlist := range IPs {
			for _, IP := range IPlist {
				fmt.Printf("%s\t%s\t%s\n", SN, IPtype, IP)
			}
		}
	}
}

//SetIPtoSNMap creates a map which the key is IP and the value is SN
func SetIPtoSNMap(oneValue wantValue) {

	for _, priIP := range oneValue.PrivateAddress {
		if net.ParseIP(priIP) == nil {
			continue
		}
		if _, ok := IPToSNMap[priIP]; ok {
			log.Printf("%s has existed\n", priIP)
			continue
		}
		IPToSNMap[priIP] = oneValue.SN
	}
	for _, pubIP := range oneValue.PublicAddress {
		if net.ParseIP(pubIP) == nil {
			continue
		}
		if _, ok := IPToSNMap[pubIP]; ok {
			log.Printf("%s has existed\n", pubIP)
			continue
		}
		IPToSNMap[pubIP] = oneValue.SN
	}
}

//SetSNtoIPMap creates a map which the key is SN and the value is IPmap
func SetSNtoIPMap(oneValue wantValue) {
	for _, priIP := range oneValue.PrivateAddress {
		privIP := SNToipMap[oneValue.SN]
		if privIP == nil {
			privIP = make(map[string][]string)
			SNToipMap[oneValue.SN] = privIP
		}
		if net.ParseIP(priIP) != nil {
			privIP["1"] = append(privIP["1"], priIP)
		}
	}
	for _, pubIP := range oneValue.PublicAddress {
		publicIP := SNToipMap[oneValue.SN]
		if publicIP == nil {
			publicIP = make(map[string][]string)
			SNToipMap[oneValue.SN] = publicIP
		}

		if net.ParseIP(pubIP) != nil {
			publicIP["2"] = append(publicIP["2"], pubIP)
		}
	}
}
