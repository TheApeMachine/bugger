package bugger

type StackType uint

const (
	FUNCTION StackType = iota
)

type Analyzer interface {
	PickupTrace(chan []byte)
}

func NewAnalyzer(analyzerType Analyzer) Analyzer {
	return analyzerType
}

type ProtoAnalyzer struct {
	stack map[StackType][]byte
}

