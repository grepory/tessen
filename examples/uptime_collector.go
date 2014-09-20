package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/grepory/tessen/metrics"
)

// StringToInt converts a string to a base-10, 32-bit int.
func StringToInt(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// GetUptime parses output from OS X's `uptime` and returns the number of
// seconds since the machine booted.
func GetUptime() float64 {
	out, err := exec.Command("uptime").Output()
	if err != nil {
		log.Fatal(err)
	}

	pieces := strings.Fields(string(out))
	days := StringToInt(pieces[2])

	var hhmm []string
	hhmm = strings.Split(pieces[4], ":")

	hours := StringToInt(hhmm[0]) + (days * 24)
	minutes := strings.Trim(hhmm[1], ",")

	uptimeString := fmt.Sprintf("%dh%sm", hours, minutes)
	uptime, err := time.ParseDuration(uptimeString)
	if err != nil {
		log.Fatal(err)
	}

	return uptime.Seconds()
}

// Collect is acn implementation of tessen.Collector.Collect() that returns the
// uptime for the system being monitored.
func Collect() []metrics.Metric {
	collected := make([]metrics.Metric, 1)

	uptime := GetUptime()

	hostname, _ := os.Hostname()

	collected[0] = metrics.Metric{
		Name:      hostname,
		Value:     uptime,
		Timestamp: time.Now().Unix(),
	}

	return collected
}

func main() {
	collector := metrics.NewCollector(Collect)
	fmt.Printf(collector.Collect())
}
