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

func main() {

	urlAddr := flag.String("u", "", "The url to ping, you can enter it without the flag")
	continuos := flag.Bool("t", false, "Continuos Ping")
	nPings := flag.Int64("n", 4, "Number of pings")
	useGet := flag.Bool("g", false, "Use GET method instead of HEAD")

	flag.Parse()

	if *urlAddr == "" {
		for i := range os.Args {
			if strings.HasPrefix(os.Args[i], "http://") || strings.HasPrefix(os.Args[i], "https://") {
				// Found!
				*urlAddr = os.Args[i]
				break
			}
		}
	}
	if *urlAddr == "" {
		fmt.Println("An Url must be provided")
		flag.Usage()
		os.Exit(1)

	}

	u, err := url.Parse(*urlAddr)
	if err != nil {
		fmt.Println("Provided Url is invalid")
		flag.Usage()
		os.Exit(1)
	}

	addr, err := net.LookupIP(u.Host)
	ip := "0.0.0.0"
	if err == nil {
		ip = addr[0].String()
	} else {
		fmt.Println("Provided Url cannot be resolved")
		flag.Usage()
		os.Exit(1)
	}

	n := int64(0)

	req := request{url: *urlAddr, logs: []logEntry{}, ip: ip, host: u.Host}

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

		tmpLog.timeEnd = time.Now().UnixMilli()
		if err != nil {
			tmpLog.statusCode = 404
			tmpLog.status = "404 Not Found"

		} else {
			tmpLog.statusCode = resp.StatusCode
			tmpLog.status = resp.Status

		}
		tmpLog.isOk = tmpLog.statusCode >= 200 && tmpLog.statusCode < 300

		req.logs = append(req.logs, tmpLog)

		printLog(tmpLog, n, req.ip, resp.ContentLength)
		n++
		if n >= *nPings && !*continuos {
			break
		}
	}

	printStats(req)
}

func printLog(l logEntry, seq int64, ip string, size int64) {

	fmt.Printf("connected to %s (%d bytes), seq=%d time=%d ms ", ip, size, seq, l.timeEnd-l.timeStart)
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
