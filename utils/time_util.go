package utils

import "time"

// Now2SecondTimestamp 当前时间转时间戳，单位：秒
func Now2SecondsTimestamp() (timestamp int64) {
	return time.Now().Unix()
}

// Now2MilliTimestamp 当前时间转时间戳，单位：毫秒
func Now2MilliTimestamp() (timestamp int64) {
	return time.Now().Unix() * 1000
}


