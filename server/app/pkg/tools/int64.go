package tools

import "strconv"

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

func Uint64ToString(e uint64) string {
	return strconv.FormatUint(e, 10)
}
