package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/grepory/tessen/metric"
)

type UptimeReader struct {
}

// ConvertSeconds converts a string to a base-10, 32-bit int.
func ConvertSeconds(s string) float64 {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// GetUptime parses output from OS X's `sysctl -n kern.boottime` and returns
// the number of seconds since the machine booted.
func GetUptime() float64 {
	out, err := exec.Command("sysctl", "-n", "kern.boottime").Output()
	// { sec = 1413585004, usec = 0 } Fri Oct 17 15:30:04 2014
	if err != nil {
		log.Fatal(err)
	}

	pieces := strings.Fields(string(out))

	return ConvertSeconds(strings.TrimRight(pieces[3], ","))
}

func (r *UptimeReader) Read() []metric.Metric {
	collected := make([]metric.Metric, 1)

	uptime := GetUptime()

	hostname, _ := os.Hostname()

	collected[0] = metric.Metric{
		Name:      hostname,
		Value:     uptime,
		Timestamp: time.Now().Unix(),
	}

	return collected
}
