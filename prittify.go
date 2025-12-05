package nlog

import (
	"github.com/hokaccha/go-prettyjson"
)

func Prettify(data map[string]any) string {
	f := prettyjson.NewFormatter()
	f.Indent = 2

	bytes, err := f.Marshal(data)
	if err != nil {
		return "{}"
	}

	return string(bytes)
}
