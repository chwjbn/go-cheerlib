package cheerlib

import "syscall"

func Process_GetNewStartProcAttr() syscall.SysProcAttr  {

	xSysProcAttr:=syscall.SysProcAttr{}

	xSysProcAttr.HideWindow=true

	return xSysProcAttr
}
