package parse

import "strconv"

// TryFuncs are the default TryFunc functions to call
// when parsing a string.
var TryFuncs = []TryFunc{
	TryInt32,
	TryInt64,
	TryUint32,
	TryUint64,
	TryFloat32,
	TryFloat64,
	TryBool,
}

// Parser represents a type capable of parsing strings.
type Parser struct {
	TryFuncs []TryFunc
}

// DefaultParser is the default Parser used by the
// String method.
var DefaultParser = New()

// String parses the string trying all TryFuncs in the
// DefaultParser.
func String(s string) interface{} {
	return DefaultParser.String(s)
}

// New creates a new Parser.
func New() *Parser {
	return &Parser{TryFuncs: TryFuncs}
}

func (p *Parser) String(s string) interface{} {
	for _, try := range p.TryFuncs {
		if v, ok := try(s); ok {
			return v
		}
	}
	return s
}

// TryFunc represents the a function capable of trying to
// convert from a string.
// TryFunc should return the cast value and true if successful,
// otherwise nil, false.
type TryFunc func(string) (interface{}, bool)

// TryInt32 tries to convert the string and returns the
// cast value if true, otherwise nil and false.
func TryInt32(s string) (interface{}, bool) {
	if v, err := strconv.ParseInt(s, 10, 32); err == nil {
		return int(v), true
	}
	return nil, false
}

// TryInt64 tries to convert the string and returns the
// cast value if true, otherwise nil and false.
func TryInt64(s string) (interface{}, bool) {
	if v, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(v), true
	}
	return nil, false
}

// TryUint32 tries to convert the string and returns the
// cast value if true, otherwise nil and false.
func TryUint32(s string) (interface{}, bool) {
	if v, err := strconv.ParseUint(s, 10, 32); err == nil {
		return int(v), true
	}
	return nil, false
}

// TryUint64 tries to convert the string and returns the
// cast value if true, otherwise nil and false.
func TryUint64(s string) (interface{}, bool) {
	if v, err := strconv.ParseUint(s, 10, 64); err == nil {
		return int(v), true
	}
	return nil, false
}

// TryFloat32 tries to convert the string and returns the
// cast value if true, otherwise nil and false.
func TryFloat32(s string) (interface{}, bool) {
	if v, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(v), true
	}
	return nil, false
}

// TryFloat64 tries to convert the string and returns the
// cast value if true, otherwise nil and false.
func TryFloat64(s string) (interface{}, bool) {
	if v, err := strconv.ParseFloat(s, 64); err == nil {
		return float32(v), true
	}
	return nil, false
}

// TryBool tries to convert the string and returns the
// cast value if true, otherwise nil and false.
func TryBool(s string) (interface{}, bool) {
	if v, err := strconv.ParseBool(s); err == nil {
		return v, true
	}
	return nil, false
}
