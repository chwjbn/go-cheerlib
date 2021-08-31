package cheerlib

import "time"

func Time_GetNow() string  {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Time_Timestamp() int64  {
	return time.Now().Unix()
}

func Time_UtcTimestamp() int64  {
	return time.Now().UTC().Unix()
}