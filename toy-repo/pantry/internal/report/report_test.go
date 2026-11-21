package report

import (
	"testing"

	"example.com/pantry/internal/item"
)

func TestAmount(t *testing.T) {
	tests := []struct {
		it   item.Item
		want string
	}{
		{item.Item{Name: "eggs", Quantity: 3, Unit: item.Piece}, "3"},
		{item.Item{Name: "flour", Quantity: 500, Unit: item.Gram}, "500g"},
		{item.Item{Name: "rice", Quantity: 2000, Unit: item.Gram}, "2kg"},
		{item.Item{Name: "sugar", Quantity: 1500, Unit: item.Gram}, "1500g"},
		{item.Item{Name: "milk", Quantity: 250, Unit: item.Millilitre}, "250ml"},
		{item.Item{Name: "oil", Quantity: 1000, Unit: item.Millilitre}, "1l"},
	}
	for _, tt := range tests {
		if got := Amount(tt.it); got != tt.want {
			t.Errorf("Amount(%+v) = %q, want %q", tt.it, got, tt.want)
		}
	}
}

func TestTable(t *testing.T) {
	items := []item.Item{
		{Name: "eggs", Quantity: 3, Unit: item.Piece},
		{Name: "flour", Quantity: 500, Unit: item.Gram},
	}
	want := "ITEM   AMOUNT\n" +
		"eggs   3\n" +
		"flour  500g\n"
	if got := Table(items); got != want {
		t.Errorf("Table() =\n%s\nwant\n%s", got, want)
	}
}

func TestTableEmpty(t *testing.T) {
	if got, want := Table(nil), "The pantry is empty.\n"; got != want {
		t.Errorf("Table(nil) = %q, want %q", got, want)
	}
}

func TestSummary(t *testing.T) {
	one := []item.Item{{Name: "eggs", Quantity: 3, Unit: item.Piece}}
	two := append(one, item.Item{Name: "flour", Quantity: 500, Unit: item.Gram})

	tests := []struct {
		items []item.Item
		want  string
	}{
		{nil, "0 items"},
		{one, "1 item"},
		{two, "2 items"},
	}
	for _, tt := range tests {
		if got := Summary(tt.items); got != tt.want {
			t.Errorf("Summary() = %q, want %q", got, tt.want)
		}
	}
}
