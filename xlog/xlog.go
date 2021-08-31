package xlog

import (
	"anynode/app/lib/cheerlib"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

type XLog struct {
	mLogChan chan string
	mLogWriter *os.File
}

func (this *XLog) recordLog(logData string)  {

	if this.mLogChan==nil{
		return
	}

	this.mLogChan<-logData

}

func (this *XLog)getNowTime() string  {
	return time.Now().Format("2006-01-02 15:04:05")
}


func (this *XLog)checkLogFile()  {

	logFilePath:=path.Join(cheerlib.Application_BaseDirectory(),"log")
	if !cheerlib.Directory_Exists(logFilePath){
		cheerlib.Directory_CreateDirectory(logFilePath)
	}

	if !cheerlib.Directory_Exists(logFilePath){
		os.Stdin.WriteString(fmt.Sprintf("[%s]XLog.checkLogFile Error For Directory_Exists(%s)=false",this.getNowTime(),logFilePath))
		return
	}

	logTime:=time.Now().Format("20060102")
	logFilePath=path.Join(logFilePath,fmt.Sprintf("xlog_%s_%s.log",cheerlib.Application_FileName(),logTime))

	//如果不存在
	if !cheerlib.File_Exists(logFilePath){

		if this.mLogWriter!=nil{
			this.mLogWriter.Close()
			this.mLogWriter=nil
		}

	}


	var xOpenFileErr error
	if this.mLogWriter==nil{

		this.mLogWriter,xOpenFileErr=os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR| os.O_APPEND, 0666)
		if xOpenFileErr!=nil{
			this.mLogWriter=nil
			os.Stdin.WriteString(fmt.Sprintf("[%s]XLog.checkLogFile xOpenFileErr=%s",this.getNowTime(),xOpenFileErr.Error()))
		}
	}

}

func (this *XLog) flushLog(logData string)  {

	logData="["+this.getNowTime()+"]"+logData+"\n"

	this.checkLogFile()

	if this.mLogWriter!=nil{

		this.mLogWriter.WriteString(logData)

	}

	os.Stdout.WriteString(logData)
}

func (this *XLog) logDaemon()  {

	for  {
		select {
		case logData:=<-this.mLogChan:
			this.flushLog(logData)
		}
	}

}

func (this *XLog)close()  {

	close(this.mLogChan)
	this.mLogChan=nil

	this.mLogWriter.Close()
	this.mLogWriter=nil
}

func initXlog() *XLog {

	pLog:=new(XLog)
	pLog.mLogChan=make(chan string,100)
	pLog.mLogWriter=nil

	go pLog.logDaemon()

	return pLog
}


var GlobalSysLogger *XLog = initXlog()

func Log(logLevel string,logContent string)  {

	_,file,line,_ := runtime.Caller(2)
	file,_=filepath.Rel(filepath.Dir(file),file)
	logData:=fmt.Sprintf("[%s][%s:%d]%s",logLevel,file,line,logContent)

	GlobalSysLogger.recordLog(logData)

}


func Info(logContent string)  {
	Log("Info",logContent)
}

func Error(logContent string)  {
	Log("Error",logContent)
}

func Close()  {
	GlobalSysLogger.close()
}
