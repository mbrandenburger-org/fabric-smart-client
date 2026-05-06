package boom

// Boom represents the boom service (Task 18).
type Boom struct{}

// New returns a new Boom service.
func New() *Boom {
	return &Boom{}
}

// Explode returns the boom message.
func (b *Boom) Explode() string {
	return "BOOM 💥"
}
