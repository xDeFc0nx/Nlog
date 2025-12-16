package nlog

import (
	"fmt"
	"maps"
)

type Log struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Trace string `json:"trace"`
}

func Info(msg string, context ...map[string]any) {
	logEntry := Log{
		Level: "INFO",
		Msg:   msg,
	}

	data := map[string]any{
		"level": logEntry.Level,
		"msg":   logEntry.Msg,
	}

	if len(context) > 0 && context[0] != nil {
		maps.Copy(data, context[0])
	}

	fmt.Println(Prettify(data))
}

func Warn(msg string, context ...map[string]any) {
	logEntry := Log{
		Level: "WARN",
		Msg:   msg,
	}

	data := map[string]any{
		"level": logEntry.Level,
		"msg":   logEntry.Msg,
	}

	if len(context) > 0 && context[0] != nil {
		maps.Copy(data, context[0])
	}

	fmt.Println(Prettify(data))
}

func Error(err error, context ...map[string]any) {
	if err == nil {
		return
	}

	file, fn, line := GetTrace(2)
	traceString := fmt.Sprintf("%s:%d  %s", file, line, fn)

	logEntry := Log{
		Level: "ERROR",
		Msg:   err.Error(),
		Trace: traceString,
	}

	data := map[string]any{
		"level": logEntry.Level,
		"msg":   logEntry.Msg,
		"trace": logEntry.Trace,
	}

	if len(context) > 0 && context[0] != nil {
		maps.Copy(data, context[0])
	}

	fmt.Println(Prettify(data))
}