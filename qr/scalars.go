package qr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// Integer accepts both numeric and quoted integers found in official QR examples.
type Integer int

func (i *Integer) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		*i = 0
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		b = []byte(s)
	}
	n, err := strconv.Atoi(string(b))
	if err != nil {
		return fmt.Errorf("qr: integer: %w", err)
	}
	*i = Integer(n)
	return nil
}

// Boolean accepts JSON booleans and the 0/1 representation used by QR.
type Boolean bool

func (v *Boolean) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case "true", "1", "\"1\"":
		*v = true
	case "false", "0", "\"0\"", "null":
		*v = false
	default:
		return fmt.Errorf("qr: invalid boolean")
	}
	return nil
}
