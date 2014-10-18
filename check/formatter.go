package check

import "fmt"

type Formatter interface {
	Format(Check) string
}

type BasicSensuFormatter struct {
}

func (f BasicSensuFormatter) Format(c Check) string {
	output := fmt.Sprintf("%s %s", c.Name, c.Status.ToString())

	if c.Message != "" {
		output = fmt.Sprintf("%s: %s", output, c.Message)
	}

	return output
}

func DefaultFormatter() Formatter {
	return BasicSensuFormatter{}
}
