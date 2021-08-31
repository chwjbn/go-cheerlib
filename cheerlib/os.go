package cheerlib

import (
	"runtime"
	"strings"
)

func Os_IsWindows() bool  {


	bRet:=false

	xOsName:=runtime.GOOS
	xOsName=strings.ToLower(xOsName)

	if strings.Contains(xOsName,"window"){
		bRet=true
	}

	return bRet

}
