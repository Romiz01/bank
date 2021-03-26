package card

import (
	"bank/pkg/bank/types"

	"fmt"
)

func ExamplePaymentSource(cards []types.Card) []types.PaymentSource {
	card := types.Card{Balance: 10_000_00, Active: true}
	a := PaymentSources(card, "card", '5058 0000 0000 8888', 10_000_00)
	fmt.Fprintln(a)

}
