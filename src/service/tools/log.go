package tools

import (
	log "github.com/cihub/seelog"
)

func init()  {
	defer log.Flush()
	// 加载日志格式
	logger, err := log.LoggerFromConfigAsFile("log.xml")
	if err != nil {
		log.Critical("err parsing config log file", err)
		return
	}

	log.ReplaceLogger(logger)
}

// 一般日志
func LogInfo(args...interface{})  {
	log.Info(args)
}

// 错误日志
func LogError(args...interface{})  {
	log.Error(args)
}

// 调试日志
func Debug(args...interface{})  {
	log.Debug(args)
}
