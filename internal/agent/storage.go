package agent

type MemStruct struct {
	metricType string
	value      string
}
type MemStorage map[string]MemStruct

func NewStorage() (memStorage MemStorage) {
	memStorage = make(map[string]MemStruct)
	return
}
