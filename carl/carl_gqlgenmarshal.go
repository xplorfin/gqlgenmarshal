// Code generated by "gqlgenmarshal -type=Carl -checkstringlowercase"; DO NOT EDIT.

package carl

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Herp-0]
	_ = x[Derp-1]
	_ = x[Merp-2]
}

func (i Carl) MarshalGQL(w io.Writer) {
	if err := json.NewEncoder(w).Encode(i.String()); err != nil {
		panic(fmt.Errorf("error marshaling Carl"))
	}
}

func (i *Carl) UnmarshalGQL(v interface{}) error {
	check, ok := v.(string)
	if !ok {
		return fmt.Errorf("Carl must be a valid string value")
	}

	switch strings.ToLower(check) {
	case strings.ToLower(Herp.String()):
		*i = Herp
	case strings.ToLower(Derp.String()):
		*i = Derp
	case strings.ToLower(Merp.String()):
		*i = Merp
	default:
		return fmt.Errorf("unknown Carl value %[1]s", check)
	}

	return nil
}
