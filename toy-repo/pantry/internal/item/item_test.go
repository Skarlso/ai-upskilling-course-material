package item

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in   string
		want Item
	}{
		{"3 eggs", Item{Name: "eggs", Quantity: 3, Unit: Piece}},
		{"500g flour", Item{Name: "flour", Quantity: 500, Unit: Gram}},
		{"2kg rice", Item{Name: "rice", Quantity: 2000, Unit: Gram}},
		{"250ml milk", Item{Name: "milk", Quantity: 250, Unit: Millilitre}},
		{"1L Olive  Oil", Item{Name: "olive oil", Quantity: 1000, Unit: Millilitre}},
		{"  12   Free Range Eggs ", Item{Name: "free range eggs", Quantity: 12, Unit: Piece}},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := Parse(tt.in)
			if err != nil {
				t.Fatalf("Parse(%q) returned error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("Parse(%q) = %+v, want %+v", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		in   string
		want error
	}{
		{"", ErrBadFormat},
		{"eggs", ErrBadFormat},
		{"three eggs", ErrBadFormat},
		{"-1 eggs", ErrBadFormat},
		{"0 eggs", ErrBadQuantity},
		{"5oz butter", ErrBadUnit},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			_, err := Parse(tt.in)
			if !errors.Is(err, tt.want) {
				t.Errorf("Parse(%q) error = %v, want %v", tt.in, err, tt.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		it   Item
		want error
	}{
		{"ok", Item{Name: "eggs", Quantity: 1, Unit: Piece}, nil},
		{"empty name", Item{Name: "   ", Quantity: 1, Unit: Piece}, ErrEmptyName},
		{"zero quantity", Item{Name: "eggs", Quantity: 0, Unit: Piece}, ErrBadQuantity},
		{"unknown unit", Item{Name: "eggs", Quantity: 1, Unit: "dozen"}, ErrBadUnit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.it.Validate(); !errors.Is(err, tt.want) {
				t.Errorf("Validate() = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		it   Item
		want string
	}{
		{Item{Name: "eggs", Quantity: 3, Unit: Piece}, "3 eggs"},
		{Item{Name: "flour", Quantity: 500, Unit: Gram}, "500g flour"},
		{Item{Name: "milk", Quantity: 250, Unit: Millilitre}, "250ml milk"},
	}
	for _, tt := range tests {
		if got := tt.it.String(); got != tt.want {
			t.Errorf("String() = %q, want %q", got, tt.want)
		}
	}
}

func TestStringRoundTrip(t *testing.T) {
	for _, in := range []string{"3 eggs", "500g flour", "250ml milk"} {
		it, err := Parse(in)
		if err != nil {
			t.Fatalf("Parse(%q): %v", in, err)
		}
		if got := it.String(); got != in {
			t.Errorf("round trip of %q gave %q", in, got)
		}
	}
}
