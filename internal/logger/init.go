/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package logger

func InitLog(loglvl string) *Slogger {
	var logger *Slogger
	switch loglvl {
	case "info":
		logger = NewJSONLogger(LevelInfo)
	case "debug":
		logger = NewJSONLogger(LevelDebug)
	case "worn":
		logger = NewJSONLogger(LevelWarn)
	}
	return logger
}
