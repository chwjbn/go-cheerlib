package cheerlib

import (
	"fmt"
	"math"
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

func Os_FormatArgs(runArg string)[]string  {

	xRunArgs:=[]string{}

	xRunArgSplit:=strings.Split(runArg," ")

	xArgTemp:=""
	xInArg:=false

	for _,xArgItem:=range xRunArgSplit{

		xArgItemVal:= strings.TrimSpace(xArgItem)
		if len(xArgItemVal)<1{
			continue
		}

		if strings.HasPrefix(xArgItemVal,"\""){
			xArgItemVal=strings.Trim(xArgItemVal,"\"")
			xArgTemp=xArgItemVal
			xInArg=true
			continue
		}

		if strings.HasSuffix(xArgItemVal,"\""){
			xArgItemVal=strings.Trim(xArgItemVal,"\"")
			xArgTemp=fmt.Sprintf("%s %s",xArgTemp,xArgItemVal)
			xRunArgs=append(xRunArgs,xArgTemp)
			xArgTemp=""
			xInArg=false
			continue
		}

		if xInArg{
			xArgTemp=fmt.Sprintf("%s %s",xArgTemp,xArgItemVal)
			continue
		}

		xRunArgs=append(xRunArgs,xArgItemVal)
	}

	return xRunArgs

}

func Os_FormatDeviceSize(data uint64,flagGiga bool) string  {

	xData:=""

	xUnit:="B"
	xBase:=1000

	if flagGiga{
		xUnit="iB"
		xBase=1024
	}


	xIndex:=1

	fData:=float64(data)
	fBase:=float64(xBase)


	xIndex=1
	if fData<math.Pow(fBase,float64(xIndex)){
		xData=fmt.Sprintf("%.2f%s",fData/math.Pow(fBase,float64(xIndex-1)),xUnit)
		return xData
	}

	xIndex=2
	if fData<math.Pow(fBase,float64(xIndex)){
		xData=fmt.Sprintf("%.2fK%s",fData/math.Pow(fBase,float64(xIndex-1)),xUnit)
		return xData
	}

	xIndex=3
	if fData<math.Pow(fBase,float64(xIndex)){
		xData=fmt.Sprintf("%.2fM%s",fData/math.Pow(fBase,float64(xIndex-1)),xUnit)
		return xData
	}

	xIndex=4
	if fData<math.Pow(fBase,float64(xIndex)){
		xData=fmt.Sprintf("%.2fG%s",fData/math.Pow(fBase,float64(xIndex-1)),xUnit)
		return xData
	}

	xIndex=5
	if fData<math.Pow(fBase,float64(xIndex)){
		xData=fmt.Sprintf("%.2fT%s",fData/math.Pow(fBase,float64(xIndex-1)),xUnit)
		return xData
	}

	xIndex=6
	if fData<math.Pow(fBase,float64(xIndex)){
		xData=fmt.Sprintf("%.2fP%s",fData/math.Pow(fBase,float64(xIndex-1)),xUnit)
		return xData
	}

	xIndex=7
	if fData<math.Pow(fBase,float64(xIndex)){
		xData=fmt.Sprintf("%.2fE%s",fData/math.Pow(fBase,float64(xIndex-1)),xUnit)
		return xData
	}

	return xData

}