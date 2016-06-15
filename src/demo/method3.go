package demo

import (
    "math"
)

type Point struct {
    X,Y float64
}

type Kc struct {
    Point
    Name string
}

func (p *Point) Abs() float64 {
    return math.Sqrt(p.X * p.X + p.Y * p.Y) 
}