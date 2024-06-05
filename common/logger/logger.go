package logger

import (
	"fmt"
	"go.uber.org/zap"
)

var (
	logInstance *zap.Logger
)

func init(){
	logInstance = zapLogger()
}

func zapLogger() *zap.Logger {
	logger, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(fmt.Sprintf("failed to initialize zap logger: %v", err))
	}
	return logger
}

func GetLogger() *zap.Logger{
	return logInstance
}

func Info(format string, args ...interface{}){
	logInstance.Info(fmt.Sprintf(format, args...))
}


func Debug(format string, args ...interface{}){
	logInstance.Debug(fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}){
	logInstance.Error(fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}){
	logInstance.Warn(fmt.Sprintf(format, args...))
}

func Sync(){
	logInstance.Sync()
}

