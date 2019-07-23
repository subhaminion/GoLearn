package main

import "fmt"

type Gopher struct {
	Name string
	Color string
	Type string
	Price int
}

func main() {
	gopher := Gopher{}

	gopher.setName()
	gopher.setColor()
	gopher.setType()
	gopher.setPrice()
	fmt.Println(gopher.getter())

}

func (g Gopher) getter() string{
	return g.Name + g.Color + g.Type
}

func (g *Gopher) setName() {
	g.Name = "Subham Bhattacharjee"
}

func (g *Gopher) setColor() {
	g.Color = "Blue is the brightest color"
}

func (g *Gopher) setType() {
	g.Type = "Ummmm!"
}

func (g *Gopher) setPrice() {
	g.Price = 1200000
}