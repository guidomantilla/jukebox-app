package factories

import (
	"fmt"
	"jukebox-app/examples/creational-patterns/factory-method/products"
)

func GetGun(gunType string) (products.Gun, error) {
	if gunType == "ak47" {
		return products.NewAk47(), nil
	}
	if gunType == "musket" {
		return products.NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
