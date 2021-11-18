package shopping_cart

type ShoppingCart struct {
	total int
}

// catalogue
// banana -> promotion
// catalogue[string]Promotion
// Promotion: (name, quantity, special price)

type Promotion struct {
	itemName     string
	quantity     int
	specialPrice int
}

var catalogue map[string]Promotion = map[string]Promotion{
	"Apple":  {"Apple", 3, 130},
	"Banana": {"Banana", 2, 45},
}

func (shopping_cart *ShoppingCart) getTotal() int {
	return shopping_cart.total
}

func (shopping_cart *ShoppingCart) addItem(
	itemName string,
	quantity int,
	cost int) {

	promotion := catalogue[itemName]

	if quantity >= promotion.quantity {
		specials := quantity / promotion.quantity
		regular := quantity % promotion.quantity

		specials_price := promotion.specialPrice * specials
		regular_price := cost * regular

		shopping_cart.total = specials_price + regular_price
		return
	}

	shopping_cart.total += cost * quantity

}
