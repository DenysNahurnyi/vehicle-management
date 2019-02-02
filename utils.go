package vehiclemanagement

import (
	"math/rand"
	"strconv"
)

// GenerateID generates ids, here I can import some lib, but to avoid dependencies I will just
// crete id's increamentally, DON'T USE THIS FOR PRODUCTION, it's here for the sake of testing
func (g *Generator) GenerateID() string {
	g.currentState++
	return strconv.Itoa(g.currentState)
}

// Generator is a struct that return new id every new call
type Generator struct {
	currentState int
}

// NewGenerator creates new Generator object
func NewGenerator() Generator {
	return Generator{
		currentState: rand.Intn(30),
	}
}
