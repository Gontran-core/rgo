package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSampleInt(t *testing.T) {
	x := NewInt([]int{1, 2, 3})
	assert.NotNil(t, x)
}

func TestSum(t *testing.T) {
	x := NewInt([]int{1, 2, 3})
	sum := x.Sum()
	assert.Equal(t, 6, sum)
}

func TestProd(t *testing.T) {
	x := NewInt([]int{2, 2, 3})
	prod := x.Prod()
	assert.Equal(t, 12, prod)
}

func TestLength(t *testing.T) {
	x := NewInt([]int{2, 2, 3})
	assert.Equal(t, 3, x.Length())
}

func TestMin(t *testing.T) {
	x := NewInt([]int{2, -2, 3})
	assert.Equal(t, -2, x.Min())
}
func TestMax(t *testing.T) {
	x := NewInt([]int{2, -2, 3})
	assert.Equal(t, 3, x.Max())
}

func TestArgmin(t *testing.T) {
	x := NewInt([]int{2, -2, 3})
	assert.Equal(t, 1, x.Argmin())
}
func TestArgminAll(t *testing.T) {
	x := NewInt([]int{-2, -2, 3})
	assert.Equal(t, []int{0, 1}, x.ArgminAll())
}
func TestArgmax(t *testing.T) {
	x := NewInt([]int{-2, 4, 4})
	assert.Equal(t, 1, x.Argmax())
}
func TestArgmaxAll(t *testing.T) {
	x := NewInt([]int{4, 4, 3})
	assert.Equal(t, []int{0, 1}, x.ArgmaxAll())
}
func TestMedian(t *testing.T) {
	x1 := NewInt([]int{1, 2, 3})
	assert.Equal(t, 2., x1.Median())
	x2 := NewInt([]int{1, 2})
	assert.Equal(t, 1.5, x2.Median())
}

func TestVar(t *testing.T) {
	x := NewInt([]int{4, 2, 0})
	// sqrt 4 + 4
	assert.InEpsilon(t, 8./2., x.Var(), 10e-8)
}

/*func TestCov(t *testing.T) {
	x := NewInt([]int{1, 2, 3})
	y := NewInt([]int{6, 5, 37})
	assert.Equal(t, 15.5, x.Cov(*y))
}*/

func TestEmptySample(t *testing.T) {
	x := NewInt([]int{})
	assert.Equal(t, 0, x.Sum())
	assert.Equal(t, 1, x.Prod())
	assert.Equal(t, 0., x.Mean()) // is that a correct math definition ?
	assert.Equal(t, 0., x.Var())
}
