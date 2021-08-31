package cheerlib

import (
	"os"
	"path/filepath"
)

func Application_BaseDirectory() string {

	sRet:=""

	xFilePath,xFilePathErr:=filepath.Abs(os.Args[0])
	if xFilePathErr!=nil{
		return sRet
	}

	sRet=filepath.Dir(xFilePath)

	return sRet

}

func Application_FileName() string  {

	sRet:=""

	xFilePath,xFilePathErr:=filepath.Abs(os.Args[0])
	if xFilePathErr!=nil{
		return sRet
	}

	sRet=filepath.Base(xFilePath)

	return sRet

}

func Application_FullPath() string {

	sRet:=""

	xFilePath,xFilePathErr:=filepath.Abs(os.Args[0])
	if xFilePathErr!=nil{
		return sRet
	}

	sRet=xFilePath

	return sRet

}
