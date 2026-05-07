package power

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPower_Activate(t *testing.T) {
	p := New(100)
	assert.Equal(t, "Power activated at level 100!", p.Activate())
}
