package slf

import (
	"strings"
	"fmt"
)

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
	Panic
	Fatal
)

func (l *Level) MarshalJSON() ([]byte, error) {
	return nil, nil // TODO
}

func (l *Level) UnmarshalJSON(data []byte) error {
	s := string(data)
	switch strings.ToLower(s) {
	case "debug":
		*l = Debug
	case "info":
		*l = Info
	case "warn":
		*l = Warn
	case "error":
		*l = Error
	case "panic":
		*l = Panic
	case "fatal":
		*l = Fatal
	default:
		return fmt.Errorf("slf: unknown level %v", s)
	}
	return nil
}