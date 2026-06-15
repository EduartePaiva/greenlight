package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// MarshalJSON implements the [encoding/json.Marshaler] interface.
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

// [encoding/json.Unmarshaler]
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	nStr, hasSuffix := strings.CutSuffix(unquotedJSONValue, " mins")
	if !hasSuffix {
		return ErrInvalidRuntimeFormat
	}

	n, err := strconv.ParseInt(nStr, 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(n)

	return nil
}
