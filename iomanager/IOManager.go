package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteDataAsJson(data any) error
}
