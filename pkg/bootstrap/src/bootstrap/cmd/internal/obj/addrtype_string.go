// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /home/ubuntu/go_1.11.11/go/src/cmd/internal/obj/addrtype_string.go

//line /home/ubuntu/go_1.11.11/go/src/cmd/internal/obj/addrtype_string.go:1
// Code generated by "stringer -type AddrType"; DO NOT EDIT.

package obj

import "strconv"

const _AddrType_name = "TYPE_NONETYPE_BRANCHTYPE_TEXTSIZETYPE_MEMTYPE_CONSTTYPE_FCONSTTYPE_SCONSTTYPE_REGTYPE_ADDRTYPE_SHIFTTYPE_REGREGTYPE_REGREG2TYPE_INDIRTYPE_REGLIST"

var _AddrType_index = [...]uint8{0, 9, 20, 33, 41, 51, 62, 73, 81, 90, 100, 111, 123, 133, 145}

func (i AddrType) String() string {
	if i >= AddrType(len(_AddrType_index)-1) {
		return "AddrType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AddrType_name[_AddrType_index[i]:_AddrType_index[i+1]]
}
