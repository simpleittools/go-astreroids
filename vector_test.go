package main

import (
	"math"
	"testing"
)

func TestVectorNormalize(t *testing.T) {
	tests := []struct {
		name     string
		vector   Vector
		expected Vector
	}{
		{
			name:     "horizontal",
			vector:   Vector{X: 3, Y: 0},
			expected: Vector{X: 1, Y: 0},
		},
		{
			name:     "diagonal",
			vector:   Vector{X: 3, Y: 4},
			expected: Vector{X: 0.6, Y: 0.8},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.vector.Normalize()

			if math.Abs(got.X-tc.expected.X) > 1e-9 {
				t.Fatalf("expected X to be %f, got %f", tc.expected.X, got.X)
			}

			if math.Abs(got.Y-tc.expected.Y) > 1e-9 {
				t.Fatalf("expected Y to be %f, got %f", tc.expected.Y, got.Y)
			}

			magnitude := math.Sqrt(got.X*got.X + got.Y*got.Y)
			if math.Abs(magnitude-1) > 1e-9 {
				t.Fatalf("expected normalized vector to have magnitude 1, got %f", magnitude)
			}
		})
	}
}
