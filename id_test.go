package gbfs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var idUnmarshalTests = []struct {
	in       string
	expected string
}{
	{"460", "460"},     // float64
	{"\"460\"", "460"}, // string
}

func TestUnmarshalID(t *testing.T) {
	var out ID

	for _, tt := range idUnmarshalTests {
		if err := json.Unmarshal([]byte(tt.in), &out); err != nil {
			t.Errorf("Failed to unmarshal JSON: %v", err)
		}

		assert.Equal(t, tt.expected, string(out))
	}
}
