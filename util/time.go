package util

import "time"

const TIME_FORMAT_ONE = "2006年01月02日 15:04:05"
const TIME_FORMAT_TWO = "2006/01/02 15:04:05"
const TIME_FORMAT_THREE = "2006—01—02 15:04:05"
const TIME_FORMAT_FOUR = "2006.01.02 15:04:05"

//将int64整形

func TimeFormat(sec int64, nsec int64, format string) string {
	return time.Unix(sec, nsec).Format(format)

}
