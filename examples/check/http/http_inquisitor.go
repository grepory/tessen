package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grepory/tessen/check"
)

type HTTPInquisitor struct {
	Target string
}

// TODO: Make this actually good... possibly even configurable and tested.
// For now we just need a working, basic implementation.
func (h HTTPInquisitor) Run() (check.Status, string, error) {
	resp, err := http.Get(h.Target)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return check.UNKNOWN, "An error has occurred", err
	}

	switch resp.StatusCode {
	case 200:
		return check.OK, buildMessage(200, h.Target), nil
	case 404:
		return check.CRITICAL, buildMessage(400, h.Target), nil
	default:
		return check.UNKNOWN, buildMessage(-1, h.Target), nil
	}
}

var statusMap = map[int]string{
	-1:  `Unknown error`,
	200: `Found`,
	404: `Not Found`,
}

func buildMessage(code int, target string) string {
	return fmt.Sprintf("%d - %s - %s", code, statusMap[code], target)
}
