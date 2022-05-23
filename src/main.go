package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main()  {
	logger()
	rt := gin.Default()
	RouterConfig(rt)
	rt.Run(":8082")
	log.Println("server start")
}

func RouterConfig(rt *gin.Engine) {
	//这个是静态资源目录
	rt.Static("/static","static")
	//gin.Recovery()相当于spring里面的操作失败退回服务注解吧
	rt.Use(gin.Recovery(),gin.Logger())
}

// 设置日志的输出位置
func logger(){
	var (
		logFileName = flag.String("log", "log/cServer.log", "Log file name")
	)
	
	runtime.GOMAXPROCS(runtime.NumCPU())
    flag.Parse()
 
    //set logfile Stdout
    logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if logErr != nil {
        fmt.Println("Fail to find", *logFile, "cServer start Failed")
        os.Exit(1)
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
 
    //write log
    log.Printf("Server abort! Cause:%v \n", "日志启动了")

}