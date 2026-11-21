package store

import (
	"errors"
	"testing"

	"example.com/pantry/internal/item"
)

func TestRenameMovesTheItem(t *testing.T) {
	s := New()
	mustAdd(t, s, "3 eggs")

	if err := s.Rename("eggs", "Duck  Eggs"); err != nil {
		t.Fatalf("Rename: %v", err)
	}
	if _, ok := s.Get("eggs"); ok {
		t.Error("the old name is still there")
	}
	got, ok := s.Get("duck eggs")
	if !ok {
		t.Fatal("the new name is missing")
	}
	if got.Name != "duck eggs" || got.Quantity != 3 || got.Unit != item.Piece {
		t.Errorf("renamed item = %+v", got)
	}
	if s.Len() != 1 {
		t.Errorf("Len() = %d, want 1", s.Len())
	}
}

func TestRenameUnknownItem(t *testing.T) {
	s := New()
	mustAdd(t, s, "3 eggs")

	if err := s.Rename("milk", "oat milk"); !errors.Is(err, ErrNotFound) {
		t.Errorf("Rename error = %v, want ErrNotFound", err)
	}
}

func TestRenameOntoExistingItemMerges(t *testing.T) {
	s := New()
	mustAdd(t, s, "500g flour")
	mustAdd(t, s, "250g plain flour")

	if err := s.Rename("plain flour", "flour"); err != nil {
		t.Fatalf("Rename: %v", err)
	}
	got, _ := s.Get("flour")
	if got.Quantity != 750 {
		t.Errorf("quantity = %d, want 750", got.Quantity)
	}
	if s.Len() != 1 {
		t.Errorf("Len() = %d, want 1", s.Len())
	}
}

func TestRenameOntoDifferentUnitFails(t *testing.T) {
	s := New()
	mustAdd(t, s, "3 eggs")
	mustAdd(t, s, "500g flour")

	if err := s.Rename("eggs", "flour"); !errors.Is(err, ErrUnitMismatch) {
		t.Fatalf("Rename error = %v, want ErrUnitMismatch", err)
	}
	eggs, ok := s.Get("eggs")
	if !ok || eggs.Quantity != 3 {
		t.Errorf("eggs changed after a failed rename: %+v (present: %v)", eggs, ok)
	}
	if flour, _ := s.Get("flour"); flour.Quantity != 500 {
		t.Errorf("flour changed after a failed rename: %+v", flour)
	}
}

func TestRenameToTheSameNameChangesNothing(t *testing.T) {
	s := New()
	mustAdd(t, s, "3 eggs")

	if err := s.Rename("eggs", " EGGS "); err != nil {
		t.Fatalf("Rename: %v", err)
	}
	if got, _ := s.Get("eggs"); got.Quantity != 3 {
		t.Errorf("quantity = %d, want 3", got.Quantity)
	}
}

func TestRenameToEmptyNameFails(t *testing.T) {
	s := New()
	mustAdd(t, s, "3 eggs")

	if err := s.Rename("eggs", "   "); !errors.Is(err, item.ErrEmptyName) {
		t.Errorf("Rename error = %v, want ErrEmptyName", err)
	}
	if _, ok := s.Get("eggs"); !ok {
		t.Error("eggs disappeared after a failed rename")
	}
}
