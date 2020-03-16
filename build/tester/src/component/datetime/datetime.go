package datetime

import (
	"regexp"
	"strconv"
)

type TimeDuration struct {
	String string
	Number float64
	Dimension string
	Seconds float64
}

func (duration *TimeDuration) InitFromString() (result *TimeDuration) {
	timeDurationRegexp := regexp.MustCompile(`^(\d+(\.\d+)?)(ns|us|ms|s)`)
	matches := timeDurationRegexp.FindStringSubmatch(duration.String)
	duration.Dimension = matches[3]
	duration.Number, _ = strconv.ParseFloat(matches[1], 64)
	switch duration.Dimension {
		case "ns":
			duration.Seconds = duration.Number / 1000000000
			break
		case "us":
			duration.Seconds = duration.Number / 1000000
			break
		case "ms":
			duration.Seconds = duration.Number / 1000
			break
		case "s":
			duration.Seconds = duration.Number
			break
	}
	return duration
}