package day5

import "github.com/samber/lo"

type RangeMap struct {
	Ranges []Range
}

func (r *RangeMap) Contains(value int) bool {
	return lo.SomeBy(r.Ranges, func(r Range) bool {
		return r.Contains(value)
	})
}

func (r *RangeMap) Map(value int) int {
	for _, r := range r.Ranges {
		if mappedValue, ok := r.Map(value); ok {
			return mappedValue
		}
	}

	return value
}
