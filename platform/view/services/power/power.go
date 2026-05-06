package power

import (
	"fmt"
)

// Power represents a power management service (Task 5).
type Power struct {
	level int
}

// New returns a new Power service.
func New(level int) *Power {
	return &Power{level: level}
}

// Activate returns a power activation message.
func (p *Power) Activate() string {
	return fmt.Sprintf("Power activated at level %d!", p.level)
}
