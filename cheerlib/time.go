package cheerlib

import "time"

func Time_GetNow() string  {

	timeFmt:="2006-01-02 15:04:05"
	return time.Now().Format(timeFmt)
}

func Time_GetTime(t time.Time) string  {
	timeFmt:="2006-01-02 15:04:05"
	return t.Format(timeFmt)
}

func Time_StrToTime(timeStr string) time.Time  {

	timeFmt:="2006-01-02 15:04:05"

	timeData,timeErr:=time.ParseInLocation(timeFmt,timeStr,time.Local)

	if timeErr!=nil{
		timeData=time.Unix(0,0)
	}

	return timeData
}

func Time_Timestamp() int64  {
	return time.Now().Unix()
}

func Time_UtcTimestamp() int64  {
	return time.Now().UTC().Unix()
}