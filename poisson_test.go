package rgo

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewPoisson(t *testing.T) {
    var x = NewPoisson(6)
    assert.Equal(t, 6., x.lambda)
}
func TestNewStdPoisson(t *testing.T) {
    var x = NewStdPoisson()
    assert.Equal(t, 1., x.lambda)
}

func TestPoissonR(t *testing.T){
    var x = NewStdPoisson()
    var x1 = x.R()
    assert.NotNil(t,x1)
    assert.True(t, x1>=0)
}
func TestPoissonRn(t *testing.T){
    var x = NewStdPoisson()
    var x1 = x.Rn(20)
    assert.Equal(t,20, len(x1))
    t.Logf("Generated values with lambda=1 : %v", x1)
}
func TestPoissonD(t *testing.T){
    var x = NewStdPoisson()
    var v1 = x.D(1)
    var v2 = x.D(5)
    var v3 = x.D(10)
    var eps = 10e-8;
    // values given by dpois(value, 10) in R
    assert.InDelta(t, 0.3678794, v1, eps)
    assert.InDelta(t, 0.003065662, v2, eps)
    assert.InDelta(t, 1.013777e-07, v3, eps)
}
func TestPoisson10D(t *testing.T){
    var x = NewPoisson(50)
    var v1 = x.D(0)
    var v2 = x.D(10)
    var v3 = x.D(25)
    var v4 = x.D(50)
    var v5 = x.D(75)
    var v6 = x.D(100)
    var eps = 10e-8;
    assert.InDelta(t, 1.92875e-22, v1, eps)
    assert.InDelta(t, 5.190544e-12, v2, eps)
    assert.InDelta(t, 3.705786e-05, v3, eps)
    assert.InDelta(t, 0.05632501, v4, eps)
    assert.InDelta(t,  0.0002057854, v5, eps)
    assert.InDelta(t,  1.630319e-10, v6, eps)
}
func TestPoissonDHuge(t *testing.T){
    var x = NewStdPoisson()
    var v1 = x.D(100)
    var v2 = x.D(20000)
    // does not end
    //var v3 = x.D(50000000000)
    var eps = 10e-8;
    assert.InDelta(t, 3.720076e-44, v1, eps)
    assert.InDelta(t, 0., v2, eps)
    //assert.InDelta(t, 0., v3, eps)
}

func TestPoissonP(t *testing.T){
}
func TestPoissonQ(t *testing.T){
}
