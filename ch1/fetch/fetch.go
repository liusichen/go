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
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

var IPToSNMap = make(map[string]string)
var SNToipMap = make(map[string]map[string][]string)
var IPtoIDCName = make(map[string]string)
var IDCNametoIP = make(map[string][]string)

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
			vidcURL(url)
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

func vidcURL(url string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resq, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "vidc:%v\n", err)
	}
	vidcData, _ := ioutil.ReadAll(resq.Body)
	defer resq.Body.Close()
	for i := 1; ; i += 2 {
		idcNum := strconv.Itoa(i)
		getName := idcNum + ".name"
		getHost := idcNum + ".hosts.#.ip"
		nameValue := gjson.Get(string(vidcData), getName)
		hostValue := gjson.Get(string(vidcData), getHost)

		if nameValue.Index == 0 {
			break
		}
		setIPtoIDC(nameValue, hostValue)
		setIDCtoIPs(nameValue, hostValue)
		for ip, name := range IPtoIDCName {
			fmt.Printf("%s\t%s\n", ip, name)
		}
		for name, ips := range IDCNametoIP {
			for _, ip := range ips {
				fmt.Printf("%s\t%s\n", name, ip)
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

func setIPtoIDC(nameValue, hostValue gjson.Result) {
	name := nameValue.String()
	for _, ip := range hostValue.Array() {
		IPtoIDCName[ip.String()] = name
	}
}

func setIDCtoIPs(nameValue, hostValue gjson.Result) {
	name := nameValue.String()
	for _, ip := range hostValue.Array() {
		IDCNametoIP[name] = append(IDCNametoIP[name], ip.String())
	}
}
