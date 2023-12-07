package day6

type Race struct {
	Time     int
	Distance int
}

func (r *Race) WaysToWin() int {
	var waysToWin int
	for i := 0; i <= r.Time; i++ {
		if i*(r.Time-i) > r.Distance {
			waysToWin++
		}
	}

	return waysToWin
}
