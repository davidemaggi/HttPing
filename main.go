package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type request struct {
	url  string
	ip   string
	host string
	logs []logEntry
}

type logEntry struct {
	status     string
	statusCode int
	timeStart  int64
	timeEnd    int64
	isOk       bool
}

/*
A simple cli util to test http connection
*/

const APP_NAME = "HttPing"
const APP_VERSION = "1.0.0"
const APP_AUTHOR = "Davide Maggi"
const APP_URL = "https://github.com/davidemaggi/HttPing"

func main() {

	urlAddr := flag.String("u", "", "The url to ping, you can enter it without the flag")
	continuos := flag.Bool("t", false, "Continuos Ping")
	nPings := flag.Int64("n", 4, "Number of pings")
	useGet := flag.Bool("g", false, "Use GET method instead of HEAD")
	showVersion := flag.Bool("v", false, "Show Version details")

	flag.Parse()

	//fmt.Printf("%#v\n", os.Args)

	if *showVersion {
		fmt.Printf("%s\n", APP_VERSION)
		//fmt.Printf("%s @: %s", APP_AUTHOR, APP_URL)
		os.Exit(0)
	}

	// If an url flag is not provided we'll check if it has been passed without flag
	if *urlAddr == "" {

		lastpar := os.Args[len(os.Args)-1]

		if strings.HasPrefix(lastpar, "http://") || strings.HasPrefix(lastpar, "https://") {
			// Found!
			*urlAddr = lastpar

		} else {

			*urlAddr = fmt.Sprintf("https://%s", lastpar)

		}

	}

	//An Url is needed... Otherwise, we'll fail
	if *urlAddr == "" {
		fmt.Println("An Url must be provided")
		flag.Usage()
		os.Exit(1)

	}

	//Let's see if it's a valid Url
	u, err := url.Parse(*urlAddr)
	if err != nil {
		fmt.Println("Provided Url is invalid")
		flag.Usage()
		os.Exit(1)
	}
	ip := "0.0.0.0"
	// Check if the user provided an url or an hostname
	tmpAddr := net.ParseIP(u.Host)
	if tmpAddr != nil {
		ip = u.Host
	} else {
		//Let's check if we can resolve host
		addr, err := net.LookupIP(u.Host)
		if err == nil {
			ip = addr[0].String()
		} else {
			fmt.Println("Provided Url cannot be resolved")
			os.Exit(1)
		}
	}

	n := int64(0)

	req := request{url: *urlAddr, logs: []logEntry{}, ip: ip, host: u.Host}

	//As default we'll just make an HEAD request
	method := ""
	if *useGet {
		method = "GET"
	} else {
		method = "HEAD"
	}

	fmt.Printf("PING: %s %s (%s): \n", u.Host, method, u.Path)

	for {
		tmpLog := logEntry{timeStart: time.Now().UnixMilli()}

		var resp *http.Response
		var err error

		if *useGet {
			resp, err = http.Get(*urlAddr)
		} else {
			resp, err = http.Head(*urlAddr)
		}

		cLength := int64(0)
		tmpLog.timeEnd = time.Now().UnixMilli()
		if err != nil {
			tmpLog.statusCode = 404
			tmpLog.status = err.Error()
			cLength = -1

		} else {
			tmpLog.statusCode = resp.StatusCode
			tmpLog.status = resp.Status
			cLength = resp.ContentLength

		}
		tmpLog.isOk = tmpLog.statusCode >= 200 && tmpLog.statusCode < 300

		req.logs = append(req.logs, tmpLog)

		printLog(tmpLog, n, req.ip, cLength)
		n++
		if n >= *nPings && !*continuos {
			break
		}
	}

	printStats(req)
}

func printLog(l logEntry, seq int64, ip string, size int64) {

	fmt.Printf("connected to %s (%d bytes), seq=%d time=%d ms : ", ip, size, seq, l.timeEnd-l.timeStart)
	fmt.Printf("%s \n", l.status)
}

func printStats(r request) {

	fmt.Printf("--- %s ping statistics ---\n", r.url)

	nAll := float64(len(r.logs))
	nOk := float64(0)
	sumTime := int64(0)

	minTime := float64(99999999999)
	maxTime := float64(0)

	for i := range r.logs {
		if r.logs[i].isOk == true {
			nOk++
		}
		curTime := r.logs[i].timeEnd - r.logs[i].timeStart
		sumTime += curTime
		if minTime > float64(curTime) {
			minTime = float64(curTime)
		}
		if maxTime < float64(curTime) {
			maxTime = float64(curTime)
		}

	}

	perc := 100 * (nAll - nOk) / nAll

	fmt.Printf("%d connects, %d ok, %.2f%% failed, time %d ms\n", int64(nAll), int64(nOk), perc, sumTime)

	avgTime := float64(sumTime) / nAll

	fmt.Printf("round-trip min/avg/max = %.1f/%.1f/%.1f ms\n", minTime, avgTime, maxTime)

}
