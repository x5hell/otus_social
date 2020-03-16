package wrk

import (
	"bufio"
	"component/datetime"
	"component/handler"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"test/helper"
)

const DefaultParameterValueThreads = 1
const DefaultParameterValueTimeout = "10s"
const DefaultParameterValueDuration = "20s"

type ParamsWrk struct {
	Threads int
	Connections int
}

type ThroughputStatistic struct {
	Average string
	StandardDeviation string
	Max string
	DeviationPercents string
}

type LatencyStatistic struct {
	Average datetime.TimeDuration
	StandardDeviation datetime.TimeDuration
	Max datetime.TimeDuration
	DeviationPercents string
}

type ErrorsWrk struct {
	Connect int
	Read int
	Write int
	Timeout int
}

type RequestStatistic struct {
	Requests int
	Duration datetime.TimeDuration
	ContentSize string
}

type StatisticWrk struct {
	Params ParamsWrk
	Latency LatencyStatistic
	Throughput ThroughputStatistic
	Request RequestStatistic
	Errors ErrorsWrk
	ThroughputPerSecond float64
	ContentPerSecond string
}

type RunWrkParameters struct {
	Url string
	Connections int
	Script string
}

type LatencyPerConnection struct {
	Latency              float64
	ConnectionsPerSecond float64
	TimeoutRate          float64
}

func RunWrk(
	url string,
	threads int,
	connections int,
	duration string,
	script string,
	timeout string,
) (result StatisticWrk, err error) {
	cmd := exec.Command(
		"wrk",
		"-t" + strconv.Itoa(threads),
		"-c" + strconv.Itoa(connections),
		"-d" + duration,
		"-s" + script,
		"--timeout",
		timeout,
		url,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		handler.ErrorLog(err)
		return result, err
	}
	wrkResponse := string(out)
	statistic := ParseWrkStat(wrkResponse)

	fmt.Println(cmd.String())
	fmt.Println(wrkResponse)

	return statistic, nil
}

func ParseWrkStat(wrkResult string) (statistic StatisticWrk) {
	wrkResultReader := strings.NewReader(wrkResult)
	scan := bufio.NewScanner(wrkResultReader)

	paramsRegexp := regexp.MustCompile(`(\d+) threads and (\d+) connections`)
	latencyRegexp := regexp.MustCompile(`Latency\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)`)
	throughputStatisticRegexp := regexp.MustCompile(`Req/Sec\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)`)
	requestRegexp := regexp.MustCompile(`(\d+) requests in (\S+), (\S+) read`)
	errorsRegexp := regexp.MustCompile(`Socket errors: connect (\d+), read (\d+), write (\d+), timeout (\d+)`)
	throughputPerSecondRegexp := regexp.MustCompile(`Requests/sec:\s+(\S+)`)
	contentPerSecondRegexp := regexp.MustCompile(`Transfer/sec:\s+(\S+)`)

	for scan.Scan() {
		switch true {
			case paramsRegexp.MatchString(scan.Text()):
				statistic.Params = ParseParams(scan.Text(), paramsRegexp)
				break
			case latencyRegexp.MatchString(scan.Text()):
				statistic.Latency = ParseLatency(scan.Text(), latencyRegexp)
				break
			case throughputStatisticRegexp.MatchString(scan.Text()):
				statistic.Throughput = ParseThroughputStatistic(scan.Text(), throughputStatisticRegexp)
				break
			case requestRegexp.MatchString(scan.Text()):
				statistic.Request = ParseRequest(scan.Text(), requestRegexp)
				break
			case errorsRegexp.MatchString(scan.Text()):
				statistic.Errors = ParseErrors(scan.Text(), errorsRegexp)
				break
			case throughputPerSecondRegexp.MatchString(scan.Text()):
				statistic.ThroughputPerSecond = ParseThroughputPerSecond(scan.Text(), throughputPerSecondRegexp)
				break
			case contentPerSecondRegexp.MatchString(scan.Text()):
				statistic.ContentPerSecond = ParseContentPerSecond(scan.Text(), contentPerSecondRegexp)
				break
		}
	}

	return statistic
}

func ParseParams(params string, regexp *regexp.Regexp) (result ParamsWrk) {
	matches := regexp.FindStringSubmatch(params)
	result.Threads, _ = strconv.Atoi(matches[1])
	result.Connections, _ = strconv.Atoi(matches[2])
	return result
}

func ParseLatency(latency string, regexp *regexp.Regexp) (result LatencyStatistic) {
	matches := regexp.FindStringSubmatch(latency)
	result.Average.String = matches[1]
	result.Average.InitFromString()
	result.StandardDeviation.String = matches[2]
	result.StandardDeviation.InitFromString()
	result.Max.String = matches[3]
	result.Max.InitFromString()
	result.DeviationPercents = matches[4]
	return result
}

func ParseThroughputStatistic(throughputStatistic string, regexp *regexp.Regexp) (result ThroughputStatistic) {
	matches := regexp.FindStringSubmatch(throughputStatistic)
	result.Average = matches[1]
	result.StandardDeviation = matches[2]
	result.Max = matches[3]
	result.DeviationPercents = matches[4]
	return result
}

func ParseRequest(request string, regexp *regexp.Regexp) (result RequestStatistic) {
	matches := regexp.FindStringSubmatch(request)
	result.Requests, _ = strconv.Atoi(matches[1])
	result.Duration.String = matches[2]
	result.Duration.InitFromString()
	result.ContentSize = matches[3]
	return result
}

func ParseErrors(errors string, regexp *regexp.Regexp) (result ErrorsWrk) {
	matches := regexp.FindStringSubmatch(errors)
	result.Connect, _ = strconv.Atoi(matches[1])
	result.Read, _ = strconv.Atoi(matches[2])
	result.Write, _ = strconv.Atoi(matches[3])
	result.Timeout, _ = strconv.Atoi(matches[4])
	return result
}

func ParseThroughputPerSecond(throughputPerSecond string, regexp *regexp.Regexp) (result float64) {
	matches := regexp.FindStringSubmatch(throughputPerSecond)
	result, _ = strconv.ParseFloat(matches[1], 64)
	return result
}

func ParseContentPerSecond(contentPerSecond string, regexp *regexp.Regexp) (result string) {
	matches := regexp.FindStringSubmatch(contentPerSecond)
	return matches[1]
}

func RunWrkUrlWithDefaultParameters(relativeUrl string, connections int, script string) (result StatisticWrk, err error) {
	fullUrl := helper.GetSocialIp() + relativeUrl
	return RunWrk(
		fullUrl, 
		DefaultParameterValueThreads, 
		connections, 
		DefaultParameterValueDuration,
		script,
		DefaultParameterValueTimeout, 
	)
}

func StatisticWrkListToLatencyPerConnection(statisticList []StatisticWrk) (result []LatencyPerConnection) {
	for _, statistic := range statisticList {
		timeoutPercents := float64(statistic.Errors.Timeout) / float64(statistic.Request.Requests)
		result = append(
			result,
			LatencyPerConnection{
				Latency:              statistic.Latency.Average.Seconds,
				ConnectionsPerSecond: statistic.ThroughputPerSecond,
				TimeoutRate:          timeoutPercents,
			},
		)
	}
	return result
}
