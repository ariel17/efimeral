package sessions

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionDistribution(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		mustFail bool
	}{
		{"valid for ubuntu", `{"distribution":"ubuntu"}`, false},
		{"valid for arch", `{"distribution":"arch"}`, false},
		{"invalid for mint", `{"distribution":"mint"}`, true},
		{"invalid JSON", `{"distribution":"`, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var d SessionDistribution
			err := json.Unmarshal([]byte(tc.input), &d)
			if tc.mustFail {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
