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

	if quantity == 3 && itemName == "Apple" {
		shopping_cart.total += 130
		return
	}

	if itemName == "Banana" && quantity >= 2 {
	    specials := quantity / 2
	    regular := quantity % 2

	    specials_price := 45 * specials
	    regular_price := cost * regular

	    shopping_cart.total = specials_price+regular_price
	    return
	}

	shopping_cart.total += cost * quantity

}
