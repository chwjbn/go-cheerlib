package cheerlib

import (
	"encoding/json"
	"golang.org/x/text/encoding"
)

func Text_StructFromJson(data interface{},dataJson string) error  {

	var xErr error=nil
	xErr=json.Unmarshal([]byte(dataJson),data)
	return xErr

}

func Text_StructToJson(data interface{}) string  {

	xData:="{}"

	jonData,jsonErr:= json.Marshal(data)

	if jsonErr!=nil{
		return xData
	}

	xData=string(jonData)

	return xData
}

func Text_GetString(data []byte,encoding encoding.Encoding) string  {

	sData:=""

	dataBuffer,xErr:=encoding.NewDecoder().Bytes(data)

	if xErr!=nil{
		return sData
	}

	sData=string(dataBuffer)

	return sData

}
