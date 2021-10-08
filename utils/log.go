/**
 * @Auth: Nuts
 * @Date: 2021/7/13 3:34 下午
 */
package utils

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func LoggerInput(method string, input interface{}) {
	params := M{
		"method": method,
		"input":  JsonEncode(input),
	}
	log.Info(params)
}

func LoggerOutput(method string, input, output interface{}, err error, start time.Time) {
	params := M{
		"method": method,
		"input":  JsonEncode(input),
		"output": JsonEncode(output),
		"times":  time.Now().Sub(start).Seconds(),
	}
	if err != nil {
		params["error"] = err.Error()
		log.Error(params)
	} else {
		log.Info(params)
	}
}

func LoggerDebug(method string, info interface{}) {
	params := M{
		"method": method,
		"input":  JsonEncode(info),
	}
	log.Debug(params)
}
