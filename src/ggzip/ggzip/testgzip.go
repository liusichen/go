package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"time"
)

var (
	cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	zipfile    = flag.String("zipfile", "ziplog", "need to zip")
	zipedfile  = flag.String("zipedfile", "zipedfile", "has zipped")
)

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Println(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	var i int
	t1 := time.Now()

	data, _ := ioutil.ReadFile(*zipfile)
	dst, _ := os.Create(*zipedfile)
	write := gzip.NewWriter(dst)
	defer write.Close()
	datalen := len(data)

	for i = 100000; i < datalen; i += 100000 {
		tmp := data[i-100000 : i]
		nBytes, err := write.Write(tmp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Write: %v\n", err)
			os.Exit(1)
		}
		afterwriteT := time.Since(t1)
		fmt.Printf("Bytes = %d\n", nBytes)
		fmt.Println("After write:", afterwriteT)
		write.Flush()
		afterflushT := time.Since(t1)
		fmt.Println("Afer flush:", afterflushT)
		time.Sleep(1e9)
	}
	tmp := data[i-100000 : datalen-1]
	nBytes, _ := write.Write(tmp)
	fmt.Printf("Bytes = %d\n", nBytes)
	time.Sleep(1e6)
	err := write.Flush()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Flush: %v\n", err)
		os.Exit(1)
	}
}
