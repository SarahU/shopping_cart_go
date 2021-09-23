package shopping_cart

type shoppingcart struct {
    total int
}

func (shopping_cart *shoppingcart) getTotal() int{
    return shopping_cart.total
}

func (shopping_cart *shoppingcart) addItem(
    itemName string,
    quantity int,
    cost int){

    shopping_cart.total = cost
}