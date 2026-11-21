// Package report turns pantry items into text for the terminal.
package report

import (
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"

	"example.com/pantry/internal/item"
)

// Amount writes a quantity for people: whole kilograms and litres are shown
// as "2kg" and "1l", everything else in the stored unit.
func Amount(it item.Item) string {
	switch it.Unit {
	case item.Gram:
		if it.Quantity%1000 == 0 {
			return strconv.Itoa(it.Quantity/1000) + "kg"
		}
		return strconv.Itoa(it.Quantity) + "g"
	case item.Millilitre:
		if it.Quantity%1000 == 0 {
			return strconv.Itoa(it.Quantity/1000) + "l"
		}
		return strconv.Itoa(it.Quantity) + "ml"
	}
	return strconv.Itoa(it.Quantity)
}

// Table lists the items in two aligned columns.
func Table(items []item.Item) string {
	if len(items) == 0 {
		return "The pantry is empty.\n"
	}
	var b strings.Builder
	w := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ITEM\tAMOUNT")
	for _, it := range items {
		fmt.Fprintf(w, "%s\t%s\n", it.Name, Amount(it))
	}
	w.Flush()
	return b.String()
}

// Summary is a one-line count of what is in the pantry.
func Summary(items []item.Item) string {
	switch len(items) {
	case 0:
		return "0 items"
	case 1:
		return "1 item"
	}
	return fmt.Sprintf("%d items", len(items))
}
