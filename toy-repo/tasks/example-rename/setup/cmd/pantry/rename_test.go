package main

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestRenameCommand(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	pantry(t, file, "add", "3", "eggs")

	code, stdout, stderr := pantry(t, file, "rename", "eggs", "duck eggs")
	if code != 0 {
		t.Fatalf("rename exited %d: %s", code, stderr)
	}
	if want := "Renamed.\n"; stdout != want {
		t.Errorf("rename printed %q, want %q", stdout, want)
	}

	_, stdout, _ = pantry(t, file, "list")
	if !strings.Contains(stdout, "duck eggs  3\n") {
		t.Errorf("list after rename printed\n%s", stdout)
	}
	if strings.Contains(stdout, "\neggs ") {
		t.Errorf("the old name is still listed\n%s", stdout)
	}
}

func TestRenameCommandUnknownItem(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	pantry(t, file, "add", "3", "eggs")

	code, _, stderr := pantry(t, file, "rename", "milk", "oat milk")
	if code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
	if !strings.Contains(stderr, "no such item") {
		t.Errorf("stderr = %q", stderr)
	}
}

func TestRenameCommandNeedsTwoNames(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	pantry(t, file, "add", "3", "eggs")

	if code, _, _ := pantry(t, file, "rename", "eggs"); code != 2 {
		t.Errorf("exit code = %d, want 2", code)
	}
}

func TestUsageMentionsRename(t *testing.T) {
	file := filepath.Join(t.TempDir(), "pantry.json")
	_, _, stderr := pantry(t, file)
	if !strings.Contains(stderr, "rename") {
		t.Errorf("usage does not mention rename:\n%s", stderr)
	}
}
