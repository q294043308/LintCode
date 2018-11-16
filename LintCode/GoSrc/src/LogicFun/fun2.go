package LogicFun

import (
	"sort"
)

// 1207. Teemo Attacking        Captain Teemo is on standby
func FindPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	sort.Ints(timeSeries)
	lastSerie := timeSeries[0]
	curSerie := timeSeries[0]
	res := 0

	for i := 1; i <= len(timeSeries); i++ {
		if i == len(timeSeries) || curSerie+duration < timeSeries[i] {
			res += curSerie - lastSerie + duration

			if i == len(timeSeries) {
				break
			}

			curSerie = timeSeries[i]
			lastSerie = timeSeries[i]
		} else {
			curSerie = timeSeries[i]
		}
	}
	return res
}
