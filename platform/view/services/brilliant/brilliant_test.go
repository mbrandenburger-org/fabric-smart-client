package brilliant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrilliant_Sparkle(t *testing.T) {
	b := New("World")
	assert.Equal(t, "Sparkle, World!", b.Sparkle())
}
