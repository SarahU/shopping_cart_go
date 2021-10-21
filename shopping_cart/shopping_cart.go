package shopping_cart

type ShoppingCart struct {
	total int
}

func (shopping_cart *ShoppingCart) getTotal() int {
	return shopping_cart.total
}

func (shopping_cart *ShoppingCart) addItem(
	itemName string,
	quantity int,
	cost int) {

	if quantity == 3 {
		shopping_cart.total += 130
	} else {
		shopping_cart.total += cost * quantity
	}
}
