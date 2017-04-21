package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var logFile = flag.String("logfile", "access.log", "the log file you want to analyse")
var smallPack = flag.Int64("s", 2048, "the limitation to  the smallest packet")
var packFilter = flag.String("filter", "", "to filter by the file format, the User or the channel by used")

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*logFile)
	ErrLog(err)

	filCount, total, averate, avertt := AnalyzeLog(data)
	fmt.Printf("the filter size is %d\nthe total is %d\nthe averate is %.3fMb/s\nthe avertt is %d\n",
		filCount, total, averate, avertt)
}

func ErrLog(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "log_anaylze: %v\n", err)
		os.Exit(1)
	}
}

func AnalyzeLog(data []byte) (int, int, float64, int64) {
	var count, total int
	var aveRttSum int64
	var aveRateSum float64
	for _, line := range strings.Split(string(data), "\n") {
		match, err := regexp.MatchString(*packFilter, line)
		var tmp = strings.Split(line, "\" \"")
		ErrLog(err)
		if match && len(tmp) > 22 {
			byteSent, err := strconv.ParseInt(tmp[11], 10, 32)
			retime, err := strconv.ParseFloat(tmp[12], 64)
			tcpRtt, err := strconv.ParseInt(tmp[22], 10, 32)
			//fmt.Printf("%d\t %.5f\t %d\n",byteSent,retime,tcpRtt)
			ErrLog(err)
			if byteSent > (*smallPack)*1024 {
				aveRateSum += Calculate(byteSent, retime, tcpRtt)
				aveRttSum += tcpRtt
				count++
			}

		}
		total++
	}
	if count == 0 {
		return count, total, 0.0, 0
	} else {
		return count, total, aveRateSum / float64(count), aveRttSum / int64(count)
	}
}

func Calculate(b int64, r float64, t int64) float64 {
	if r+float64(t) == 0.0 && r == 0 {
		return 0.0
	}
	floatT := float64(t) / 1000 / 1000
	return float64(b) / (r + floatT) * 8 / 1000 / 1000
}
