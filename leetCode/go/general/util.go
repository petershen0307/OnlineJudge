package general

import (
	"runtime"
)

func GetCallerName() string {
	// the grandparent is the actual caller we interest
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName
}
