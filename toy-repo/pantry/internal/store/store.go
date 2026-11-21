// Package store keeps the pantry's items in memory and on disk.
package store

import (
	"errors"
	"fmt"
	"sort"

	"example.com/pantry/internal/item"
)

var (
	ErrNotFound     = errors.New("store: no such item")
	ErrUnitMismatch = errors.New("store: item is kept in a different unit")
	ErrInsufficient = errors.New("store: not enough in stock")
)

// Store holds items by their normalized name.
type Store struct {
	items map[string]item.Item
}

// New returns an empty store.
func New() *Store {
	return &Store{items: make(map[string]item.Item)}
}

// Add puts an item in the store. If an item with the same name is already
// there, the quantities are added together. Adding to an existing item in a
// different unit is an error, and the store is left unchanged.
func (s *Store) Add(it item.Item) error {
	it.Name = item.NormalizeName(it.Name)
	if err := it.Validate(); err != nil {
		return err
	}
	if have, ok := s.items[it.Name]; ok {
		if have.Unit != it.Unit {
			return fmt.Errorf("%w: %s is kept in %s", ErrUnitMismatch, it.Name, have.Unit)
		}
		it.Quantity += have.Quantity
	}
	s.items[it.Name] = it
	return nil
}

// Remove takes qty out of the named item, counted in the item's own unit.
// Taking out everything deletes the item.
func (s *Store) Remove(name string, qty int) error {
	name = item.NormalizeName(name)
	have, ok := s.items[name]
	if !ok {
		return fmt.Errorf("%w: %s", ErrNotFound, name)
	}
	if qty <= 0 {
		return item.ErrBadQuantity
	}
	if qty > have.Quantity {
		return fmt.Errorf("%w: only %s left", ErrInsufficient, have)
	}
	if qty == have.Quantity {
		delete(s.items, name)
		return nil
	}
	have.Quantity -= qty
	s.items[name] = have
	return nil
}

// Get returns the named item, if it is in the store.
func (s *Store) Get(name string) (item.Item, bool) {
	it, ok := s.items[item.NormalizeName(name)]
	return it, ok
}

// List returns every item, sorted by name.
func (s *Store) List() []item.Item {
	out := make([]item.Item, 0, len(s.items))
	for _, it := range s.items {
		out = append(out, it)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// Len returns the number of different items in the store.
func (s *Store) Len() int {
	return len(s.items)
}
