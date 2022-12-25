package main

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func main() {
	//r := gin.Default()  //默认使用的Logger和Recovery中间件
	//使用的日志默认是Logger(gin的记录日志的中间件)

	//func Default() *Engine {
	//	debugPrintWARNINGDefault()
	//	engine := New()
	//	engine.Use(Logger(), Recovery())  //恢复
	//	return engine
	//}

	InitLogger()
	r := gin.New() //没有自带中间件的路由(gin的引擎,http service)(可以将gin的engine作为Handler作为http.Server的参数再进一层封装)
	//gin框架集成zap(zap可以处理gin框架的日志了,其余的看喜好,zap,logrus,原生log都可以)
	r.Use(GinLogger(Logger), GinRecovery(Logger, true))

	//默认端口8080
	r.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.Run()
}

//中间件实现这些功能,自己封装的gin的中间件
//zap中间件
//GinLogger接收gin框架的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery //原始请求参数
		//上边的数据先拿到,确定好,避免被某些中间件修改request信息(比如各种网关功能的中间件)

		ctx.Next() //往后边的中间件发送,洋葱结构

		//花费的时间(请求)
		cost := time.Since(start) //时长
		//msg:请求的路径  //field key:value
		logger.Info(path,
			//指定了类型性能会更好
			//zap.Logger的结构化,如果是sugaredLogger就没必要设置了,毕竟不支持
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			//客户端的ip和机子信息或者比如postman的信息
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("error", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),

			//专门用来存放的time.Duration(time.Since返回的是time.Duration)
			zap.Duration("cost", cost), //时间间隔
		)
	}
}

//recover捕获程序的panic
//原生的gin的recovery改的
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc { //错误会记录堆栈信息
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				if stack {
					logger.Error(
						"[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error(
						"[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
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
