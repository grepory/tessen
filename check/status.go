package check

type Status int

const (
	OK       Status = iota
	WARNING  Status = iota
	CRITICAL Status = iota
	UNKNOWN  Status = iota
)

var statusLookupTable = map[Status]string{
	OK:       `OK`,
	WARNING:  `WARNING`,
	CRITICAL: `CRITICAL`,
	UNKNOWN:  `UNKNOWN`,
}

func (s Status) ToString() string {
	return statusLookupTable[s]
}

func (s Status) ToInt() int {
	return int(s)
}
