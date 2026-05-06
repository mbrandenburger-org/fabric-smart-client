package boom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoom_Explode(t *testing.T) {
	b := New()
	assert.Equal(t, "BOOM 💥", b.Explode())
}
