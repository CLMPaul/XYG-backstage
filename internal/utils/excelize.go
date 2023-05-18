package utils

import "time"

// 修正 excelize v2.4.0 要求 time.Time 必须是 UTC 时区的限制
// https://github.com/360EntSecGroup-Skylar/excelize/issues/409
func FixExcelizeTime(t time.Time) time.Time {
	t, _ = time.ParseInLocation("2006-01-02 15:04:05", t.Local().Format("2006-01-02 15:04:05"), time.UTC)
	return t
}
