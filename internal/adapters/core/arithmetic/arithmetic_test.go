package arithmetic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Addition(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int32(3), answer)
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Subtraction(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int32(-1), answer)
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Multiplication(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int32(2), answer)
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()
	answer, err := arith.Division(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, int32(0), answer)
}
