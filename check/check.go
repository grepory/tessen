package check

import (
	"fmt"
	"log"
	"os"
)

type Check struct {
	Inquisitor Inquisitor
	Formatter  Formatter
	Status     Status
	Name       string
	Message    string
}

func NewCheck(name string, i Inquisitor) Check {
	return Check{
		Inquisitor: i,
		Name:       name,
		Formatter:  DefaultFormatter(),
	}
}

func NewCheckWithFormatter(name string, i Inquisitor, f Formatter) Check {
	return Check{
		Inquisitor: i,
		Name:       name,
		Formatter:  f,
	}
}

func (c Check) Run() {
	status, message, err := c.Inquisitor.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(UNKNOWN.ToInt())
	}
	c.Status, c.Message = status, message

	fmt.Println(c.Formatter.Format(c))

	exit(c.Status)
}

func exit(s Status) {
	os.Exit(s.ToInt())
}
