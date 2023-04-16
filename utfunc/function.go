package utfunc

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/DeniesKresna/gohelper/utstring"
)

/*
	Param: StackFunction : int
	determine how many previous function layer name you want to get
	1 will return the function name call this function
	2 will return the function name call the function who call this function
	etc ..

	Result will be in this format objstruct-functionName

	TODO: Need More Test, Do not use this
*/
func GetObjectFunctionName(stackFunction int) string {
	pc, _, _, _ := runtime.Caller(1)
	nm := runtime.FuncForPC(pc).Name()
	nms := strings.Split(nm, ".")
	if len(nms) < 3 {
		return nm
	}
	var fnTempl string
	for _, v := range nms[1:] {
		v = utstring.RemoveChars(v, []string{"(", "*", ")"})
		fnTempl += fmt.Sprintf("%s-", v)
	}
	fnTempl = strings.TrimSuffix(fnTempl, "-")
	return fnTempl
}
