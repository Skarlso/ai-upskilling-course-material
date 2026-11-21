package store

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"example.com/pantry/internal/item"
)

func mustAdd(t *testing.T, s *Store, in string) {
	t.Helper()
	it, err := item.Parse(in)
	if err != nil {
		t.Fatalf("Parse(%q): %v", in, err)
	}
	if err := s.Add(it); err != nil {
		t.Fatalf("Add(%q): %v", in, err)
	}
}

func TestAddMergesQuantities(t *testing.T) {
	s := New()
	mustAdd(t, s, "500g flour")
	mustAdd(t, s, "1kg Flour")

	got, ok := s.Get("flour")
	if !ok {
		t.Fatal("flour is missing")
	}
	if got.Quantity != 1500 {
		t.Errorf("quantity = %d, want 1500", got.Quantity)
	}
	if s.Len() != 1 {
		t.Errorf("Len() = %d, want 1", s.Len())
	}
}

func TestAddUnitMismatch(t *testing.T) {
	s := New()
	mustAdd(t, s, "3 eggs")

	err := s.Add(item.Item{Name: "eggs", Quantity: 100, Unit: item.Gram})
	if !errors.Is(err, ErrUnitMismatch) {
		t.Fatalf("Add error = %v, want ErrUnitMismatch", err)
	}
	if got, _ := s.Get("eggs"); got.Quantity != 3 || got.Unit != item.Piece {
		t.Errorf("store changed after a failed Add: %+v", got)
	}
}

func TestAddRejectsInvalidItem(t *testing.T) {
	s := New()
	if err := s.Add(item.Item{Name: "eggs", Quantity: 0, Unit: item.Piece}); !errors.Is(err, item.ErrBadQuantity) {
		t.Errorf("Add error = %v, want ErrBadQuantity", err)
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}
}

func TestRemove(t *testing.T) {
	s := New()
	mustAdd(t, s, "500g flour")

	if err := s.Remove("Flour", 200); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	if got, _ := s.Get("flour"); got.Quantity != 300 {
		t.Errorf("quantity = %d, want 300", got.Quantity)
	}

	if err := s.Remove("flour", 300); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	if _, ok := s.Get("flour"); ok {
		t.Error("flour is still there after removing all of it")
	}
}

func TestRemoveErrors(t *testing.T) {
	s := New()
	mustAdd(t, s, "3 eggs")

	tests := []struct {
		name string
		item string
		qty  int
		want error
	}{
		{"not found", "milk", 1, ErrNotFound},
		{"too many", "eggs", 4, ErrInsufficient},
		{"zero", "eggs", 0, item.ErrBadQuantity},
		{"negative", "eggs", -1, item.ErrBadQuantity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Remove(tt.item, tt.qty); !errors.Is(err, tt.want) {
				t.Errorf("Remove(%q, %d) = %v, want %v", tt.item, tt.qty, err, tt.want)
			}
		})
	}
	if got, _ := s.Get("eggs"); got.Quantity != 3 {
		t.Errorf("store changed after failed removes: %+v", got)
	}
}

func TestListIsSortedByName(t *testing.T) {
	s := New()
	mustAdd(t, s, "250ml milk")
	mustAdd(t, s, "3 eggs")
	mustAdd(t, s, "500g flour")

	var names []string
	for _, it := range s.List() {
		names = append(names, it.Name)
	}
	want := []string{"eggs", "flour", "milk"}
	if len(names) != len(want) {
		t.Fatalf("List() names = %v, want %v", names, want)
	}
	for i := range want {
		if names[i] != want[i] {
			t.Fatalf("List() names = %v, want %v", names, want)
		}
	}
}

func TestSaveAndLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pantry.json")

	s := New()
	mustAdd(t, s, "3 eggs")
	mustAdd(t, s, "500g flour")
	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if loaded.Len() != 2 {
		t.Fatalf("Len() = %d, want 2", loaded.Len())
	}
	if got, _ := loaded.Get("flour"); got.Quantity != 500 || got.Unit != item.Gram {
		t.Errorf("flour = %+v", got)
	}
}

func TestLoadMissingFileIsEmpty(t *testing.T) {
	s, err := Load(filepath.Join(t.TempDir(), "nope.json"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}
}

func TestLoadRejectsBrokenFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "pantry.json")
	if err := os.WriteFile(path, []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(path); err == nil {
		t.Error("Load accepted a broken file")
	}
}
