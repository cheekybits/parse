package parse_test

import (
	"fmt"
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/parse"
)

const (
	maxint32 int64  = 2147483647
	maxint64 uint64 = 9223372036854775807
)

func TestParse(t *testing.T) {
	is := is.New(t)

	parser := parse.New()

	for _, test := range []struct {
		S string
		V interface{}
	}{
		{
			"100", int(100),
		},
		{
			fmt.Sprintf("%v", maxint32+1), int64(maxint32 + 1),
		},
		// TODO(tylerb): For some reason, this doesnt't work and I can't
		// figure out why.
		// {
		// 	fmt.Sprintf("%v", maxint64+1), uint64(maxint64 + 1),
		// },
		{
			"1.1", 1.1,
		},
		{
			"9223372036854775807.1", 9.223372e+18,
		},
		{
			"a string", "a string",
		},
		{
			"true", true,
		},
		{
			"false", false,
		},
	} {

		v := parser.String(test.S)
		is.Equal(v, test.V)

	}

}
