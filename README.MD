# Slog

## 介绍
Slog 是根据zap日志库进行二次封装的一个日志库,可以实现日志切割,
不同环境下生成不同应用名称的日志，且支持控制台+文件两种模式同时输出。


## env使用

env中如有AppName字段且值不为空，日志文件则会使用该名字进行生成.
env中如有XDEBUG字段且值不为空，日志将会在控制台输出以及文件输出.

其他使用方式请参考zap官网

## 日志使用

请参考测试用例，如需记录错误详细内容日志，请使用`ErrorLog()`对错误结果进行调用

示例日志使用代码:
```
package xxxx

import (
	"fmt"
	"github.com/iiiusky/slog"
	"go.uber.org/zap"
	"strings"
)

// 普通日志
func Logger() *zap.Logger {
	return slog.Logger(&slog.SLoggerSetting{
		AppName:    strings.ToLower(AppName),
		Path:       LogPath,
		IsDebug:    LogDebug,
		CallerSkip: 1,
	})
}

// 专门记录错误日志
func ErrorLog(errorCode int, msg string, err error) {
	Logger().With(zap.Int("ErrorCode", errorCode)).Error(msg, slog.ErrorLog(err))
}

func DebugF(msg string, args ...interface{}) {
	Logger().Debug(fmt.Sprintf(msg, args...))
}

func Debug(msg string, fields ...zap.Field) {
	Logger().Debug(msg, fields...)
}

func InfoF(msg string, args ...interface{}) {
	Logger().Info(fmt.Sprintf(msg, args...))
}

func Info(msg string, fields ...zap.Field) {
	Logger().Info(msg, fields...)
}

func WarnF(msg string, args ...interface{}) {
	Logger().Warn(fmt.Sprintf(msg, args...))
}

func Warn(msg string, fields ...zap.Field) {
	Logger().Warn(msg, fields...)
}

func ErrorF(msg string, args ...interface{}) {
	Logger().Error(fmt.Sprintf(msg, args...))
}

func Error(msg string, fields ...zap.Field) {
	Logger().Error(msg, fields...)
}

func DPanicF(msg string, args ...interface{}) {
	Logger().DPanic(fmt.Sprintf(msg, args...))
}

func DPanic(msg string, fields ...zap.Field) {
	Logger().DPanic(msg, fields...)
}

func PanicF(msg string, args ...interface{}) {
	Logger().Panic(fmt.Sprintf(msg, args...))
}

func Panic(msg string, fields ...zap.Field) {
	Logger().Panic(msg, fields...)
}

func FatalF(msg string, args ...interface{}) {
	Logger().Fatal(fmt.Sprintf(msg, args...))
}

func Fatal(msg string, fields ...zap.Field) {
	Logger().Fatal(msg, fields...)
}
```
