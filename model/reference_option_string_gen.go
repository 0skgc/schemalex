// Code generated by "stringer -type=ReferenceOption -output=reference_option_string_gen.go"; DO NOT EDIT.

package model

import "fmt"

const _ReferenceOption_name = "ReferenceOptionNoneReferenceOptionRestrictReferenceOptionCascadeReferenceOptionSetNullReferenceOptionNoAction"

var _ReferenceOption_index = [...]uint8{0, 19, 42, 64, 86, 109}

func (i ReferenceOption) String() string {
	if i < 0 || i >= ReferenceOption(len(_ReferenceOption_index)-1) {
		return fmt.Sprintf("ReferenceOption(%d)", i)
	}
	return _ReferenceOption_name[_ReferenceOption_index[i]:_ReferenceOption_index[i+1]]
}
