// Package item defines what the pantry stores and how it is written down.
package item

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Unit is how an item is counted.
type Unit string

const (
	Piece      Unit = "pc"
	Gram       Unit = "g"
	Millilitre Unit = "ml"
)

// Item is one line in the pantry.
type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Unit     Unit   `json:"unit"`
}

var (
	ErrEmptyName   = errors.New("item: name is empty")
	ErrBadQuantity = errors.New("item: quantity must be positive")
	ErrBadUnit     = errors.New("item: unknown unit")
	ErrBadFormat   = errors.New(`item: expected "<amount> <name>"`)
)

// NormalizeName lower-cases a name and collapses runs of whitespace, so that
// "Olive  Oil" and "olive oil" are the same item.
func NormalizeName(s string) string {
	return strings.Join(strings.Fields(strings.ToLower(s)), " ")
}

// Validate reports whether the item can be stored.
func (i Item) Validate() error {
	if NormalizeName(i.Name) == "" {
		return ErrEmptyName
	}
	if i.Quantity <= 0 {
		return ErrBadQuantity
	}
	switch i.Unit {
	case Piece, Gram, Millilitre:
		return nil
	}
	return fmt.Errorf("%w: %q", ErrBadUnit, i.Unit)
}

// Parse reads an item written as "<amount> <name>".
//
// The amount is a whole number with an optional unit: "3 eggs", "500g flour",
// "2kg rice", "250ml milk", "1l oil". Kilograms are stored as grams and litres
// as millilitres. An amount without a unit counts pieces.
func Parse(s string) (Item, error) {
	fields := strings.Fields(s)
	if len(fields) < 2 {
		return Item{}, ErrBadFormat
	}
	qty, unit, err := parseAmount(fields[0])
	if err != nil {
		return Item{}, err
	}
	it := Item{
		Name:     NormalizeName(strings.Join(fields[1:], " ")),
		Quantity: qty,
		Unit:     unit,
	}
	if err := it.Validate(); err != nil {
		return Item{}, err
	}
	return it, nil
}

func parseAmount(s string) (int, Unit, error) {
	digits, suffix := s, ""
	for i, r := range s {
		if r < '0' || r > '9' {
			digits, suffix = s[:i], s[i:]
			break
		}
	}
	n, err := strconv.Atoi(digits)
	if err != nil {
		return 0, "", fmt.Errorf("%w: bad amount %q", ErrBadFormat, s)
	}
	switch strings.ToLower(suffix) {
	case "":
		return n, Piece, nil
	case "g":
		return n, Gram, nil
	case "kg":
		return n * 1000, Gram, nil
	case "ml":
		return n, Millilitre, nil
	case "l":
		return n * 1000, Millilitre, nil
	}
	return 0, "", fmt.Errorf("%w: %q", ErrBadUnit, suffix)
}

// String writes the item the way Parse reads it.
func (i Item) String() string {
	if i.Unit == Piece {
		return fmt.Sprintf("%d %s", i.Quantity, i.Name)
	}
	return fmt.Sprintf("%d%s %s", i.Quantity, i.Unit, i.Name)
}
