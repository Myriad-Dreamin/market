// Code generated by "stringer -type=MethodType"; DO NOT EDIT.

package controller

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Any-0]
	_ = x[GET-1]
	_ = x[POST-2]
	_ = x[DELETE-3]
	_ = x[PATCH-4]
	_ = x[PUT-5]
	_ = x[OPTIONS-6]
	_ = x[HEAD-7]
	_ = x[Static-8]
	_ = x[StaticFile-9]
	_ = x[StaticFS-10]
}

const _MethodType_name = "AnyGETPOSTDELETEPATCHPUTOPTIONSHEADStaticStaticFileStaticFS"

var _MethodType_index = [...]uint8{0, 3, 6, 10, 16, 21, 24, 31, 35, 41, 51, 59}

func (i MethodType) String() string {
	if i >= MethodType(len(_MethodType_index)-1) {
		return "MethodType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MethodType_name[_MethodType_index[i]:_MethodType_index[i+1]]
}
