package shopping_cart

import (
	"testing"
)

func TestShoppingCart(t *testing.T) {

	t.Run("shopping cart total = 0", func(t *testing.T) {
		cart := ShoppingCart{}
		total := cart.getTotal()

		if total != 0 {
			t.Errorf("expected %d received %d", 0, total)
		}
	})

	t.Run("shopping cart total = 0", func(t *testing.T) {
		cart := ShoppingCart{0}

		cart.addItem("Apple", 1, 20)

		total := cart.getTotal()

		if total != 20 {
			t.Errorf("expected %q received %q", 0, total)
		}
	})

	t.Run("shopping cart total = 50 for 1 apple", func(t *testing.T) {
		cart := ShoppingCart{0}

		cart.addItem("Apple", 1, 50)

		total := cart.getTotal()

		if total != 50 {
			t.Errorf("expected %d received %d", 0, total)
		}
	})

	t.Run("shopping cart total = 80 for 1 apple and 1 banana", func(t *testing.T) {
		cart := ShoppingCart{0}

		cart.addItem("Apple", 1, 50)
		cart.addItem("Banana", 1, 30)

		total := cart.getTotal()

		if total != 80 {
			t.Errorf("expected %d received %d", 0, total)
		}
	})

	t.Run("shopping cart total = 130 for 2 apples and 1 banana", func(t *testing.T) {
		cart := ShoppingCart{0}

		cart.addItem("Apple", 2, 50)
		cart.addItem("Banana", 1, 30)

		total := cart.getTotal()

		if total != 130 {
			t.Errorf("expected %d received %d", 0, total)
		}
	})
}
