package day5

type Range struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

func (r *Range) Contains(value int) bool {
	return value >= r.SourceStart && value < r.SourceStart+r.Length
}

func (r *Range) Map(value int) (int, bool) {
	if r.Contains(value) {
		return value + r.DestinationStart - r.SourceStart, true
	}

	return value, false
}
