package main

import "fmt"

type geometry interface {
	area() int
	perim() int
}

type react struct {
	width, height int
}

type react666 struct {
	width, height int
}

type circle struct {
	radius int
}

func (r react ) area() int {
	return r.width * r.height
}

func (r react ) perim()  int{
	return 2*r.width + 2*r.height
}

func (c circle ) area() int {
	return c.radius * c.radius
}

func (c circle ) perim()  int{
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
