package cheerlib

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

func Encrypt_Md5(data string) string {

	xHash := md5.New()
	xHash.Write([]byte(data))

	return  hex.EncodeToString(xHash.Sum(nil))
}

func Encrypt_FileMd5(filePath string) string {

	xData:=""

	if !File_Exists(filePath){
		return xData
	}

	xFileData, xFileErr := ioutil.ReadFile(filePath)
	if xFileErr!=nil{
		return xData
	}

	xHash := md5.New()
	xData=hex.EncodeToString(xHash.Sum(xFileData))

	return  xData
}

func Encrypt_Base64Encode(data string) string  {

	xData:=base64.StdEncoding.EncodeToString([]byte(data))
	return xData

}

func Encrypt_Base64Decode(data string) string  {

	xData:=""

	xDataBuf,xDataErr:=base64.StdEncoding.DecodeString(data)

	if xDataErr!=nil{
		return xData
	}

	xData=string(xDataBuf)

	return xData

}

func Encrypt_UrlEncode(data string) string  {

	xData:=""

	xData=url.QueryEscape(data)

	return xData

}

func Encrypt_UrlDecode(data string) string  {

	xData:=""

	xDataTemp,xDataTempErr:=url.QueryUnescape(data)

	if xDataTempErr!=nil{
		return xData
	}

	xData=xDataTemp

	return xData

}


func Encrypt_NewId() string {

	xEnvStr:=strings.Join(os.Environ(),"|")
	xEnvStr=Encrypt_Md5(xEnvStr)

	xNewId:=fmt.Sprintf("%d",time.Now().Nanosecond())
	rand.Seed(time.Now().Unix())
	xNewId=fmt.Sprintf("msg-id-%s-%d",xNewId,rand.Int63())
	xNewId=Encrypt_Md5(xNewId+xEnvStr)

	return xNewId

}