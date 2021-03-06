package cheerlib

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func Process_ExcuteCommand(timeout int,command string,args ...string) ([]byte,error) {

	LogTag("Process_ExcuteCommand Begin")
	LogInfo(fmt.Sprintf("Process_ExcuteCommand timeout=[%d],command=[%s],args=[%s]",timeout,command,strings.Join(args," ")))
	defer func() {
		LogTag("Process_ExcuteCommand End")
	}()

	var xError error=nil
	var xData []byte=[]byte{}

	var xResultBuf bytes.Buffer
	var xErrorBuf bytes.Buffer

	var xCmdHandle *exec.Cmd
	xCmdHandle=exec.Command(command, args...)

	xCmdHandle.Stdout=&xResultBuf
	xCmdHandle.Stderr=&xErrorBuf
	xCmdHandle.Env=os.Environ()

	xError=xCmdHandle.Start()
	if xError!=nil{
		return xData,xError
	}

	xDoneChan := make(chan error)
	go func() {
		xDoneChan <- xCmdHandle.Wait()
	}()


	xTimeOut:=timeout
	if xTimeOut<0{
		xTimeOut=5
	}


	select {

	    case <-time.After(time.Duration(xTimeOut) * time.Second):
	    	xError=errors.New(fmt.Sprintf("Process_ExcuteCommand TimeOut After %ds",xTimeOut))
			LogError("Process_ExcuteCommand With TimeOutError:"+xError.Error())

	    	if timeout>0{
				xCmdHandle.Process.Signal(syscall.SIGINT)
				time.Sleep(50*time.Millisecond)
				xCmdHandle.Process.Kill()
			}

	    case xDoneErr:=<-xDoneChan:
	    	if xDoneErr!=nil{
				LogError("Process_ExcuteCommand With DoneError:"+xDoneErr.Error())
			}

			xError=nil
	}

	if xError!=nil{
		return xData,xError
	}

	if xErrorBuf.Len()>0{
		xError=errors.New("Process_ExcuteCommand Result Return Errors")
		xData=xErrorBuf.Bytes()
		return xData,xError
	}

	if xResultBuf.Len()>0{
		xError=nil
		xData=xResultBuf.Bytes()
	}

	return xData,xError
}

func Process_Start(command string,args ...string) (error,int) {

	LogTag("Process_Start Begin")
	LogInfo(fmt.Sprintf("Process_Start command=[%s],args=[%s]",command,strings.Join(args," ")))
	defer func() {
		LogTag("Process_Start End")
	}()


	xPid:=0

	var xCmdHandle *exec.Cmd
	xCmdHandle=exec.Command(command, args...)

	xProcAttr:=Process_GetNewStartProcAttr()

	xCmdHandle.Env=os.Environ()
	xCmdHandle.SysProcAttr=&xProcAttr

	xError:=xCmdHandle.Start()

	if xError!=nil{
		return xError,xPid
	}

	if xCmdHandle.Process!=nil{
		xPid=xCmdHandle.Process.Pid
	}

	return xError,xPid
}

func Process_StartNew(command string,args ...string) error {

	xCmdFile:="bash"
	xCmdFileArg:="-c"
	xCmdRun:=fmt.Sprintf("nohup %s >/dev/null 2>&1 &",fmt.Sprintf("%s %s",command,strings.Join(args," ")))

	if Os_IsWindows(){
		xCmdFile="cmd"
		xCmdFileArg="/c"
		xCmdRun=fmt.Sprintf("start %s",fmt.Sprintf("%s %s",command,strings.Join(args," ")))
	}


	_,xErr:=Process_ExcuteCommand(60,xCmdFile,xCmdFileArg,xCmdRun)


	return xErr

}

func Process_Kill(pid int)  {

	xProc,xProcErr:=os.FindProcess(pid)
	if xProcErr==nil{
		if xProc!=nil{
		    xProc.Kill()
		    time.Sleep(100*time.Millisecond)
		}
	}

}

func Process_KillForce(pid int)  {

	xCmdFile:="bash"
	xCmdFileArg:="-c"
	xCmdKill:=fmt.Sprintf("kill -9 %d",pid)

	if Os_IsWindows(){
		xCmdFile="cmd"
		xCmdFileArg="/c"
		xCmdKill=fmt.Sprintf("taskkill /PID %d",pid)
	}

	Process_ExcuteCommand(10,xCmdFile,xCmdFileArg,xCmdKill)

}

func Process_KillByName(processName string)  {

	xCmdFile:="bash"
	xCmdFileArg:="-c"
	xCmdKill:=fmt.Sprintf("ps -ef |grep %s |awk '{print $2}'|xargs kill -9",processName)

	if Os_IsWindows(){
		xCmdFile="cmd"
		xCmdFileArg="/c"
		xCmdKill=fmt.Sprintf("taskkill /im %d",processName)
	}

	Process_ExcuteCommand(10,xCmdFile,xCmdFileArg,xCmdKill)

}