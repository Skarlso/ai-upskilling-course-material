package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

// pantry runs the command against a pantry file and returns the exit code
// with everything it printed.
func pantry(t *testing.T, file string, args ...string) (code int, stdout, stderr string) {
	t.Helper()
	var out, errOut bytes.Buffer
	code = run(append([]string{"--file", file}, args...), &out, &errOut)
	return code, out.String(), errOut.String()
}

func TestAddThenList(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")

	if code, _, stderr := pantry(t, file, "add", "500g", "flour"); code != 0 {
		t.Fatalf("add exited %d: %s", code, stderr)
	}
	code, stdout, stderr := pantry(t, file, "add", "1kg", "Flour")
	if code != 0 {
		t.Fatalf("add exited %d: %s", code, stderr)
	}
	if want := "Added. You now have 1500g flour.\n"; stdout != want {
		t.Errorf("add printed %q, want %q", stdout, want)
	}
	if code, _, stderr := pantry(t, file, "add", "3", "eggs"); code != 0 {
		t.Fatalf("add exited %d: %s", code, stderr)
	}

	code, stdout, stderr = pantry(t, file, "list")
	if code != 0 {
		t.Fatalf("list exited %d: %s", code, stderr)
	}
	want := "ITEM   AMOUNT\n" +
		"eggs   3\n" +
		"flour  1500g\n" +
		"2 items\n"
	if stdout != want {
		t.Errorf("list printed\n%s\nwant\n%s", stdout, want)
	}
}

func TestListEmpty(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	code, stdout, _ := pantry(t, file, "list")
	if code != 0 {
		t.Fatalf("list exited %d", code)
	}
	if want := "The pantry is empty.\n0 items\n"; stdout != want {
		t.Errorf("list printed %q, want %q", stdout, want)
	}
}

func TestRemove(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	pantry(t, file, "add", "1l", "olive", "oil")

	if code, _, stderr := pantry(t, file, "remove", "olive", "oil", "400"); code != 0 {
		t.Fatalf("remove exited %d: %s", code, stderr)
	}
	_, stdout, _ := pantry(t, file, "list")
	if !strings.Contains(stdout, "olive oil  600ml\n") {
		t.Errorf("list after remove printed\n%s", stdout)
	}
}

func TestRemoveMoreThanThereIs(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	pantry(t, file, "add", "3", "eggs")

	code, _, stderr := pantry(t, file, "remove", "eggs", "5")
	if code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
	if !strings.Contains(stderr, "not enough in stock") {
		t.Errorf("stderr = %q", stderr)
	}
}

func TestBadCommandLine(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	tests := [][]string{
		{},
		{"frobnicate"},
		{"remove", "eggs"},
	}
	for _, args := range tests {
		if code, _, _ := pantry(t, file, args...); code != 2 {
			t.Errorf("pantry %v exited %d, want 2", args, code)
		}
	}
}

func TestAddRejectsNonsense(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	if code, _, _ := pantry(t, file, "add", "some", "eggs"); code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
}
