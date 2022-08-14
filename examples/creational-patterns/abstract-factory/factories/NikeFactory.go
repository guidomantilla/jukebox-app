package factories

import (
	"jukebox-app/examples/creational-patterns/abstract-factory/products"
)

type NikeFactory struct {
}

func (n *NikeFactory) MakeShoe() products.Shoe {
	return &products.NikeShoe{
		Shoe: &products.AbstractShoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *NikeFactory) MakeShirt() products.Shirt {
	return &products.NikeShirt{
		Shirt: &products.AbstractShirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
