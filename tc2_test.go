package tc2

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_tc2_1(t *testing.T) {
	tc := Init()
	assert.NotEqual(t, tc, nil)
}
