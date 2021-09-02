package shopping_cart

import ( "testing" )

func TestShoppingCart(t *testing.T) {

        t.Run("A=1", func(t *testing.T) {
            cart = shoppingcart{0}
            total = cart.getTotal()

            if(total != 0){
                t.Errorf("expected %q received %q", 0, total)
            }
        })
}