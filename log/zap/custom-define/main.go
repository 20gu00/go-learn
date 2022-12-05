package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func main() {
	//New自定义配置
	//zapcore.Core:Encoder WriterSyncer LogLevel
	//Encoder编码器,开箱即用的NewJSONEncoder
	//WriterSyncer日志写到哪里去,比如打开的文件句柄
	//日志级别类似水坝,高于水坝的可以过去,也就是高级别的日志可以通过

	InitLogger()
	Logger.Info("aaa")
	SugarLogger.Info("bbb")
}

func InitLogger() {
	writerSyncer := getLogWriter()
	logEncoder := getLogEncoder()
	//配置
	core := zapcore.NewCore(logEncoder, writerSyncer, zapcore.DebugLevel)

	//初始化Logger
	//将函数调用方的信息也打印出日志(函数的文件,行号)
	Logger = zap.New(core, zap.AddCaller())
	SugarLogger = Logger.Sugar()

}

func getLogEncoder() zapcore.Encoder {
	//使用默认的生产环境用的encoder
	//text
	//生产环境的日志不容易度 ISO8601TimeEncoder,直接设置它的返回值
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, //zapcore.EpochTimeEncoder,  1970.1.1到现在的时间戳  也可以自定义方法
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)

	//参考着两个的配置项
	//return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()) //json console mapObject
}

func getLogWriter() zapcore.WriteSyncer {
	//openFile 权限
	//每一次都创建,这样只会是一次性写入,会覆盖
	//file, _ := os.Create("./zap.log")
	//读写方式打开
	//file, _ := os.OpenFile("./zap.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	//return zapcore.AddSync(file)

	//LumberJack
	//备份文件名加了时间戳
	lumberjack := &lumberjack.Logger{
		Filename:   "./zap.log",
		MaxSize:    10, //m
		MaxBackups: 5,
		MaxAge:     30,   //天
		Compress:   true, //默认false

	}
	return zapcore.AddSync(lumberjack)
}
