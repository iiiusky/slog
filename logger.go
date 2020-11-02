/*
Copyright © 2020 iiusky sky@03sec.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package slog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type SLoggerSetting struct {
	AppName    string
	Path       string
	IsDebug    bool
	CallerSkip int
}

func Logger(userSLoggerSetting ...*SLoggerSetting) *zap.Logger {
	var sLoggerSetting *SLoggerSetting

	if len(userSLoggerSetting) == 1 {
		sLoggerSetting = userSLoggerSetting[0]
	} else {
		sLoggerSetting = &SLoggerSetting{
			AppName:    "Application",
			Path:       "./",
			IsDebug:    false,
			CallerSkip: 1,
		}
	}

	var writers []zapcore.WriteSyncer

	hook := lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", sLoggerSetting.Path, sLoggerSetting.AppName), // 日志文件路径
		MaxSize:    128,                                                                   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                                                                    // 日志文件最多保存多少个备份
		MaxAge:     7,                                                                     // 文件最多保存多少天
		Compress:   true,                                                                  // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder, // 全路径编码器
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	var allCore []zapcore.Core
	if sLoggerSetting.IsDebug {
		atomicLevel.SetLevel(zap.DebugLevel)
		// 控制台输出日志
		allCore = append(allCore, zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),                // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台
			atomicLevel))
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	writers = append(writers, zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)))
	// 普通文件记录日志
	allCore = append(allCore, zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),   // 编码器配置
		zapcore.NewMultiWriteSyncer(writers...), // 打印到文件
		atomicLevel,                             // 日志级别
	))

	// 设置输出文件名和行号
	caller := zap.AddCaller()
	// 输出的文件名和行号是调用封装函数的位置
	callerSkip := zap.AddCallerSkip(sLoggerSetting.CallerSkip)
	// 开启堆栈
	stacktrace := zap.AddStacktrace(zapcore.WarnLevel)

	// 设置初始化字段
	filed := zap.Fields(zap.String("appName", sLoggerSetting.AppName))

	core := zapcore.NewTee(allCore...)
	// 构造日志
	logger := zap.New(core, caller, filed, stacktrace, callerSkip)

	return logger
}

// 设置错误日志记录
func ErrorLog(err error) zap.Field {
	return zap.Error(err)
}
