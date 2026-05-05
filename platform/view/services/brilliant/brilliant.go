package brilliant

import (
	"fmt"
)

// Brilliant is a dummy service for the contribution process testbed.
type Brilliant struct {
	name string
}

// New returns a new Brilliant service.
func New(name string) *Brilliant {
	return &Brilliant{name: name}
}

// Sparkle returns a brilliant message.
func (b *Brilliant) Sparkle() string {
	return fmt.Sprintf("Sparkle, %s!", b.name)
}
