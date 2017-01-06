package dna

import (
	"math/rand"
	"time"
	"fmt"
)

const MUTATION_RATE = 0.01

// The structure for DNA, which is a list of genes
type DNA struct {
	genes        []float64
}

// Create a DNA sequence with a given list of genes
func CreateDNA(genes []float64) *DNA {
	g := make([]float64, len(genes))
	for i := 0; i < len(genes); i++ {
		if genes[i] < 0 {
			g[i] = 0.
		} else if genes[i] > 1 {
			g[i] = 1.
		}	else {
			g[i] = genes[i]
		}
	}

	return &DNA{g}
}

// Create a DNA seqence of length n, chosen randomly
func CreateRandDNA(n int) *DNA {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	g := make([]float64, n)
	for i := 0; i < n; i++ {
		g[i] = r.Float64()
	}
	return &DNA{g}
}

// TODO create a crossover strand
func (self *DNA) Crossover(partner *DNA) *DNA {
	return self
}

// TODO mutate a strand
func (self *DNA) Mutate() *DNA {
	return self
}

// TODO calculate the fitness of the DNA
func (self *DNA) Fitness() float64 {
	return 0.
}

// Get number of genes
func (self *DNA) Length() int {
	return len(self.genes)
}

// Print DNA properly
func (self *DNA) String() string {
	s := "["
	for i := 0; i < self.Length(); i++ {
		if i != 0 {
			s += ", "
		}
		// %.2f => float number with precision 2
		s += fmt.Sprintf("%.2f", self.genes[i])
	}
	s += "]"
	return s
}