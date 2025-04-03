package checkdigit_test

import (
	"os"
	"testing"

	"github.com/MacoTasu/go-barcode/checkdigit"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCalculateModulus10All(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected bool
	}{
		{
			name:     "valid checkdigit - last digit matches calculated checkdigit",
			digits:   []int{4, 9, 7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 7},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, checkdigit.ValidateModulus10All(3, tt.digits))
		})
	}
}
