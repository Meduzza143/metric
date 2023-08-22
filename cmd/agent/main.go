package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"

	senders "github.com/Meduzza143/metric/cmd/agent/internal"
)

type MemStruct struct {
	metricType string
	value      string
}

var MemStorage = make(map[string]MemStruct)

var pollInterval time.Duration
var reportInterval time.Duration

func uintToStr(value uint64) string {
	return strconv.FormatUint(value, 10)
}

func uint32ToStr(value uint32) string {
	return strconv.FormatUint(uint64(value), 10)
}

func poller() {
	var mem runtime.MemStats
	var i uint64 = 0
	fmt.Println("poller>>>")

	for {
		runtime.ReadMemStats(&mem)
		i++

		MemStorage["Alloc"] = MemStruct{"gauge", uintToStr(mem.Alloc)}
		MemStorage["BuckHashSys"] = MemStruct{"gauge", uintToStr(mem.BuckHashSys)}
		MemStorage["Frees"] = MemStruct{"gauge", uintToStr(mem.Frees)}
		MemStorage["GCSys"] = MemStruct{"gauge", uintToStr(mem.GCSys)}
		MemStorage["HeapAlloc"] = MemStruct{"gauge", uintToStr(mem.HeapAlloc)}
		MemStorage["HeapIdle"] = MemStruct{"gauge", uintToStr(mem.HeapIdle)}
		MemStorage["HeapInuse"] = MemStruct{"gauge", uintToStr(mem.HeapInuse)}
		MemStorage["HeapObjects"] = MemStruct{"gauge", uintToStr(mem.HeapObjects)}
		MemStorage["HeapReleased"] = MemStruct{"gauge", uintToStr(mem.HeapReleased)}
		MemStorage["HeapSys"] = MemStruct{"gauge", uintToStr(mem.HeapSys)}
		MemStorage["LastGC"] = MemStruct{"gauge", uintToStr(mem.LastGC)}
		MemStorage["Lookups"] = MemStruct{"gauge", uintToStr(mem.Lookups)}
		MemStorage["MCacheInuse"] = MemStruct{"gauge", uintToStr(mem.MCacheInuse)}
		MemStorage["MCacheSys"] = MemStruct{"gauge", uintToStr(mem.MCacheSys)}
		MemStorage["MSpanInuse"] = MemStruct{"gauge", uintToStr(mem.MSpanInuse)}

		MemStorage["MSpanSys"] = MemStruct{"gauge", uintToStr(mem.MSpanSys)}
		MemStorage["Mallocs"] = MemStruct{"gauge", uintToStr(mem.Mallocs)}
		MemStorage["NextGC"] = MemStruct{"gauge", uintToStr(mem.NextGC)}
		MemStorage["OtherSys"] = MemStruct{"gauge", uintToStr(mem.OtherSys)}
		MemStorage["PauseTotalNs"] = MemStruct{"gauge", uintToStr(mem.PauseTotalNs)}
		MemStorage["StackInuse"] = MemStruct{"gauge", uintToStr(mem.StackInuse)}
		MemStorage["StackSys"] = MemStruct{"gauge", uintToStr(mem.StackSys)}
		MemStorage["Sys"] = MemStruct{"gauge", uintToStr(mem.Sys)}
		MemStorage["TotalAlloc"] = MemStruct{"gauge", uintToStr(mem.TotalAlloc)}

		MemStorage["NumForcedGC"] = MemStruct{"gauge", uint32ToStr(mem.NumForcedGC)}
		MemStorage["NumGC"] = MemStruct{"gauge", uint32ToStr(mem.NumGC)}

		MemStorage["GCCPUFraction"] = MemStruct{"gauge", strconv.FormatFloat(mem.GCCPUFraction, 'f', -1, 64)}
		MemStorage["RandomValue"] = MemStruct{"gauge", strconv.FormatFloat(rand.Float64(), 'f', -1, 64)}

		MemStorage["PollCount"] = MemStruct{"counter", uintToStr(i)}

		time.Sleep(pollInterval)
	}
}

func sender() {
	for {
		fmt.Println("sender>>>")
		for k, v := range MemStorage {
			status := senders.SendData(v.value, k, v.metricType)
			if status != 200 {
				log.Fatalf("status: %d", status)
			}
		}
		time.Sleep(reportInterval)
	}
}

func main() {
	fmt.Println("client_start")

	flagAdrPtr := flag.String("a", "localhost:8080", "endpont address:port")
	flagRepPtr := flag.Duration("r", 10*time.Second, "report interval in seconds")
	flagPolPtr := flag.Duration("p", 2*time.Second, "poll interval in seconds")

	flag.Parse()

	adr, ok := os.LookupEnv("ADDRESS")
	if !ok {
		adr = *flagAdrPtr
	}

	repString, ok := os.LookupEnv("REPORT_INTERVAL")
	if !ok {
		reportInterval = *flagRepPtr
	} else {
		reportInterval, _ = time.ParseDuration(repString)
	}

	pollString, ok := os.LookupEnv("POLL_INTERVAL")
	if !ok {
		pollInterval = *flagPolPtr
	} else {
		pollInterval, _ = time.ParseDuration(pollString)
	}

	senders.APIURL = "http://" + adr

	go poller()
	go sender()

	for {
		time.Sleep(500 * time.Second)
	}

}
