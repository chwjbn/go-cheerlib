package cheerlib

import "syscall"

func Process_GetNewStartProcAttr() syscall.SysProcAttr  {

	xSysProcAttr:=syscall.SysProcAttr{}

	xSysProcAttr.Setsid=true
	//xSysProcAttr.Setpgid=true

	return xSysProcAttr
}
