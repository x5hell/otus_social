package wrkTest

import (
	"component/config"
	"component/graphic"
	"component/wrk"
	"fmt"
	"os"
	"time"
)

const ResultDirectory = "test" + string(os.PathSeparator) + "result" + string(os.PathSeparator)
const EnvLuaScriptsPath = "LUA_SCRIPTS_PATH"
const EnvTestSearchScript = "TEST_SEARCH_SCRIPT"

func TestUrl(testUrl string) (wrkResultList []wrk.StatisticWrk, err error) {
	wrkParametersList := TestParametersList(testUrl)
	for _, wrkParameters := range wrkParametersList {
		wrkResult, err := wrk.RunWrkUrlWithDefaultParameters(
			wrkParameters.Url,
			wrkParameters.Connections,
			wrkParameters.Script,
		)
		time.Sleep(10 * time.Second)
		if err != nil {
			return wrkResultList, err
		}
		wrkResultList = append(wrkResultList, wrkResult)
	}

	return wrkResultList, nil
}

func TestParametersList(testUrl string) (wrkParametersList []wrk.RunWrkParameters) {
	luaScriptsPath, _ := config.GetEnv(EnvLuaScriptsPath)
	testSearchScript, _ := config.GetEnv(EnvTestSearchScript)
	testSearchScriptPath := luaScriptsPath + testSearchScript

	return []wrk.RunWrkParameters{
		{
			Url: testUrl,
			Connections: 1,
			Script: testSearchScriptPath,
		},
		{
			Url: testUrl,
			Connections: 50,
			Script: testSearchScriptPath,
		},
		{
			Url: testUrl,
			Connections: 100,
			Script: testSearchScriptPath,
		},
		{
			Url: testUrl,
			Connections: 200,
			Script: testSearchScriptPath,
		},
		{
			Url: testUrl,
			Connections: 400,
			Script: testSearchScriptPath,
		},
		{
			Url: testUrl,
			Connections: 800,
			Script: testSearchScriptPath,
		},
	}
}

func TestLatencyAndThroughput(testUrl string, testPrefix string)  {
	wrkResultList, err := TestUrl(testUrl)
	if err == nil {
		drawConnectionsLatency(wrkResultList, testPrefix)
		drawConnectionsThroughput(wrkResultList, testPrefix)
	}

}

func drawConnectionsLatency(wrkResultList []wrk.StatisticWrk, testPrefix string) {
	canvas := graphic.Data2D{
		XSize: 700,
		YSize: 700,
		Data:  []graphic.FloatPoint{},
		XName: "connections",
		YName: "latency, sec / timeouts",
	}

	for _, wrkStatistic := range wrkResultList {
		canvas.Data = append(
			canvas.Data,
			graphic.FloatPoint{
				X:           float64(wrkStatistic.Params.Connections),
				Y:           wrkStatistic.Latency.Average.Seconds,
				Description: fmt.Sprint(wrkStatistic.Errors.Timeout),
			},
		)
	}

	_ = graphic.Draw2DGraphic(canvas, ResultDirectory + testPrefix + "latencyTest.png")
}

func drawConnectionsThroughput(wrkResultList []wrk.StatisticWrk, testPrefix string) {
	canvas := graphic.Data2D{
		XSize: 700,
		YSize: 700,
		Data:  []graphic.FloatPoint{},
		XName: "connections",
		YName: "throughput, req/sec",
	}

	for _, wrkStatistic := range wrkResultList {
		canvas.Data = append(
			canvas.Data,
			graphic.FloatPoint{
				X:  float64(wrkStatistic.Params.Connections),
				Y:  wrkStatistic.ThroughputPerSecond,
			},
		)
	}

	_ = graphic.Draw2DGraphic(canvas, ResultDirectory + testPrefix + "throughputTest.png")
}

