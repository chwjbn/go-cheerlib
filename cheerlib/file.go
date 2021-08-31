package cheerlib

import (
	"io"
	"io/ioutil"
	"os"
)

func File_Exists(path string) bool{

	bRet:=false

	xFileInfo,xFileInfoErr:=os.Stat(path)
	if xFileInfoErr!=nil{
		return bRet
	}

	if !xFileInfo.IsDir(){
		bRet=true
	}

	return bRet

}

func File_Delete(path string) bool{

	bRet:=false

	xErr:=os.Remove(path)
	if xErr==nil{
		bRet=true
	}

	return bRet

}

func File_Rename(oldPath string,newPath string) bool{

	bRet:=false

	xErr:=os.Rename(oldPath,newPath)
	if xErr==nil{
		bRet=true
	}

	return bRet

}

func File_Copy(srcPath string,desPath string) bool{

	bRet:=false

	if !File_Exists(srcPath){
		return bRet
	}

	if File_Exists(desPath){
		return bRet
	}

	xDesFile,xDesFileErr:=os.Create(desPath)
	if xDesFileErr!=nil{
		return bRet
	}

	defer xDesFile.Close()


	xSrcFile,xSrcFileErr:=os.Open(srcPath)

	if xSrcFileErr!=nil{
		return bRet
	}

	defer xSrcFile.Close()


	_,xCopyErr:=io.Copy(xDesFile,xSrcFile)

	if xCopyErr!=nil{
		return bRet
	}

	bRet=true

	return bRet

}

func File_ReadAllText(path string) string  {

	sData:=""

	xFileData,xFileDataErr:=ioutil.ReadFile(path)
	if xFileDataErr!=nil{
		return sData
	}

	sData=string(xFileData)

	return sData

}

func File_WriteAllText(path string,data string) bool  {

	bRet:=false

	xFileDataErr:=ioutil.WriteFile(path,[]byte(data),0644)
	if xFileDataErr==nil{
		bRet=true
	}

	return bRet

}