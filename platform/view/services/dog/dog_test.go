package dog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDog_Walk(t *testing.T) {
	d := New("Buddy")
	assert.Equal(t, "Walking the dog: Buddy. Good boy!", d.Walk())
}
