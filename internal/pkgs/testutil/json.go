package testutil

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// DeleteMapAnyByPath remove a specific field by path
// separate by dot (.) from an unstructured data.
func DeleteMapAnyByPath(mp map[string]any, path string) {
	curr := any(mp)
	paths := strings.Split(path, ".")
	for _, p := range paths[:len(paths)-1] {
		cur, ok := curr.(map[string]any)
		if !ok {
			break
		}

		next, ok := cur[p]
		if !ok {
			break
		}
		curr = next
	}
	if cur, ok := curr.(map[string]any); ok {
		delete(cur, paths[len(paths)-1])
	}
}

func AssertJsonEq(t *testing.T, expected, actual string, ignoreFields ...string) {
	var ma, mb map[string]any
	if err := json.NewDecoder(strings.NewReader(expected)).Decode(&ma); err != nil {
		panic(err)
	}
	if err := json.NewDecoder(strings.NewReader(actual)).Decode(&mb); err != nil {
		panic(err)
	}

	for _, f := range ignoreFields {
		DeleteMapAnyByPath(ma, f)
		DeleteMapAnyByPath(mb, f)
	}

	assert.Equal(t, ma, mb)
}
