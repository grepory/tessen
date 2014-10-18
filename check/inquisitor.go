package check

type Inquisitor interface {
	Run() (Status, string, error)
}
