package main

import (
	"fmt"
	"jukebox-app/examples/creational-patterns/factory-method/factories"
	"jukebox-app/examples/creational-patterns/factory-method/products"
)

func main() {
	ak47, _ := factories.GetGun("ak47")
	musket, _ := factories.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g products.Gun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
