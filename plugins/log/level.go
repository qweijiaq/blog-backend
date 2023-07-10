package log

import "encoding/json"

type Level int

const (
	DebugLevel Level = 1
	InfoLevel  Level = 2
	WarnLevel  Level = 3
	ErrorLevel Level = 4
)

func (l Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l Level) String() string {
	var str string
	switch l {
	case DebugLevel:
		str = "debug"
	case InfoLevel:
		str = "info"
	case WarnLevel:
		str = "warning"
	case ErrorLevel:
		str = "error"
	}

	return str
}
