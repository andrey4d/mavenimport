/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package logger

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Error(string, ...any)
	Warn(string, ...any)
}
