package log

import "fmt"

func Printf(format string, args ...interface{}) {
	printf(fmt.Sprintf(format, args...))
}
