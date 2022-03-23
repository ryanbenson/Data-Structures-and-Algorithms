package shortesttimeinterval

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func smallestTimeInterval(times []string, minTimeIncrementInSec int64) string {
	// set min time to be something crazy high so it'll be updated by our min time easily
	var minTime int64 = 99999999999999
	var maxTime int64 = 0

	timesMap := make(map[int64]int64)
	for _, givenTime := range times {
		timeSplit := strings.Split(givenTime, ":")
		hour, _ := strconv.Atoi(timeSplit[0])
		minute, _ := strconv.Atoi(timeSplit[1])
		t := time.Date(2022, time.November, 10, hour, minute, 0, 0, time.UTC).Unix()

		if t < minTime {
			minTime = t
		}
		if t > maxTime {
			maxTime = t
		}
		val, ok := timesMap[t]
		if ok {
			val = val + 1
		} else {
			timesMap[t] = 1
		}
	}
	sortedTimes := []int64{}
	// increment by minute
	for i := minTime; i <= maxTime; i = i + minTimeIncrementInSec {
		_, ok := timesMap[i]
		if ok {
			for timesMap[i] > 0 {
				sortedTimes = append(sortedTimes, i)
				timesMap[i]--
			}
		}
	}

	var shortestInterval int64 = 999999999999999
	totalTimes := len(sortedTimes)
	for k := 0; k < totalTimes; k++ {
		// if we're at the end, don't bother checking for the next time
		if k == totalTimes - 1 {
			break
		}
		timeDiff := sortedTimes[k + 1] - sortedTimes[k]
		if timeDiff < shortestInterval {
			shortestInterval = timeDiff
		}
	}
	
	// get how long the duration was in human friendly format
	timeStamp := time.Unix(shortestInterval, 0).UTC()
	hour, min, _ := timeStamp.Clock()
	if hour > 0 {
		return fmt.Sprintf("%d hours, %d minutes", hour, min)
	}
	return fmt.Sprintf("%d minutes", min)
}

func smallestTimeIntervalSimple(times []string) string {
	intTimes := []int{}
	for _, givenTime := range times {
		timeSplit := strings.Split(givenTime, ":")
		hour, _ := strconv.Atoi(timeSplit[0])
		minute, _ := strconv.Atoi(timeSplit[1])
		t := time.Date(2022, time.November, 10, hour, minute, 0, 0, time.UTC).Unix()
		intTimes = append(intTimes, int(t))
	}
	sort.Ints(intTimes)

	var shortestInterval int = 999999999999999
	totalTimes := len(intTimes)
	for k := 0; k < totalTimes; k++ {
		// if we're at the end, don't bother checking for the next time
		if k == totalTimes - 1 {
			break
		}
		timeDiff := intTimes[k + 1] - intTimes[k]
		if timeDiff < shortestInterval {
			shortestInterval = timeDiff
		}
	}
	
	// get how long the duration was in human friendly format
	timeStamp := time.Unix(int64(shortestInterval), 0).UTC()
	hour, min, _ := timeStamp.Clock()
	if hour > 0 {
		return fmt.Sprintf("%d hours, %d minutes", hour, min)
	}
	return fmt.Sprintf("%d minutes", min)
}
