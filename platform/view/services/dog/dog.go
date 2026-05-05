package dog

import (
	"fmt"
)

// Dog is a dummy service for the contribution process testbed (Task 12).
type Dog struct {
	name string
}

// New returns a new Dog.
func New(name string) *Dog {
	return &Dog{name: name}
}

// Walk returns a message about walking the dog.
func (d *Dog) Walk() string {
	return fmt.Sprintf("Walking the dog: %s. Good boy!", d.name)
}
