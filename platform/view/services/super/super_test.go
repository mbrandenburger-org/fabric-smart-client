package super

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuper_Blast(t *testing.T) {
	s := New(9000)
	assert.Equal(t, "Super Blast with power 9000!", s.Blast())
}
