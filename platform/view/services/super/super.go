package super

import (
	"fmt"
)

// Super is a dummy service for the contribution process testbed (Task 4).
type Super struct {
	power int
}

// New returns a new Super service.
func New(power int) *Super {
	return &Super{power: power}
}

// Blast returns a super message.
func (s *Super) Blast() string {
	return fmt.Sprintf("Super Blast with power %d!", s.power)
}
