// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /home/ubuntu/go_1.11.11/go/src/math/big/roundingmode_string.go

//line /home/ubuntu/go_1.11.11/go/src/math/big/roundingmode_string.go:1
// Code generated by "stringer -type=RoundingMode"; DO NOT EDIT.

package big

import "strconv"

const _RoundingMode_name = "ToNearestEvenToNearestAwayToZeroAwayFromZeroToNegativeInfToPositiveInf"

var _RoundingMode_index = [...]uint8{0, 13, 26, 32, 44, 57, 70}

func (i RoundingMode) String() string {
	if i >= RoundingMode(len(_RoundingMode_index)-1) {
		return "RoundingMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RoundingMode_name[_RoundingMode_index[i]:_RoundingMode_index[i+1]]
}
