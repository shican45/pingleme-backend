//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package conf

import (
	"PingLeMe-Backend/cache"
	"PingLeMe-Backend/model"
	"PingLeMe-Backend/util"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	var logLevel int
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		logLevel = util.LevelDebug
	case "info":
		logLevel = util.LevelInformational
	case "warn":
		logLevel = util.LevelWarning
	case "error":
		logLevel = util.LevelError
	default:
		logLevel = util.LevelInformational
	}

	logMaxSize, err1 := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	if err1 != nil {
		panic("env error! failed to set config.")
	}

	logMaxAge, err2 := strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	if err2 != nil {
		panic("env error! failed to set config.")
	}

	logMaxBackUp, err3 := strconv.Atoi(os.Getenv("LOG_MAX_BACKUP"))
	if err3 != nil {
		panic("env error! failed to set config.")
	}

	logCompress, err4 := strconv.ParseBool(os.Getenv("LOG_COMPRESS"))
	if err4 != nil {
		panic("env error! failed to set config.")
	}

	logJSONFormat, err5 := strconv.ParseBool(os.Getenv("LOG_JSON_FORMAT"))
	if err5 != nil {
		panic("env error! failed to set config.")
	}

	logShowLines, err6 := strconv.ParseBool(os.Getenv("LOG_SHOW_LINES"))
	if err6 != nil {
		panic("env error! failed to set config.")
	}

	logShowInConsole, err7 := strconv.ParseBool(os.Getenv("LOG_SHOW_IN_CONSOLE"))
	if err7 != nil {
		panic("env error! failed to set config.")
	}

	// 设置日志级别
	util.InitLogger(
		os.Getenv("LOG_PATH"),
		logLevel,
		logMaxSize,
		logMaxBackUp,
		logMaxAge,
		logCompress,
		logJSONFormat,
		logShowLines,
		logShowInConsole,
	)

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", zap.Error(err))
	}

	var DBLogLevel logger.LogLevel
	switch os.Getenv("DB_LOG_LEVEL") {
	case "silent":
		DBLogLevel = logger.Silent
	case "warn":
		DBLogLevel = logger.Warn
	case "error":
		DBLogLevel = logger.Error
	case "info":
		DBLogLevel = logger.Info
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"), DBLogLevel)
	cache.Redis()
}
