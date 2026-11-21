Add a way to rename an item in the pantry.

People make typos when they add things, and today the only fix is to remove the item and add it again.

What it should do:

- `Store` gets a method `Rename(oldName, newName string) error`.
- Names are normalized the same way as everywhere else in the store, so case and extra spaces don't matter.
- Renaming an item that is not in the store returns `ErrNotFound`.
- Renaming onto a name that already exists adds the two quantities together, as long as both are kept in the same unit. If the units differ it returns `ErrUnitMismatch` and nothing changes.
- Renaming an item to the name it already has changes nothing and is not an error.
- A new name that is empty returns `item.ErrEmptyName` and nothing changes.
- The command line gets `pantry rename <old> <new>`. It takes exactly two names, so a name with spaces has to be quoted. On success it prints `Renamed.` and saves the file. The usage text lists the new command.

There are failing tests for this in `internal/store/rename_test.go` and `cmd/pantry/rename_test.go`. Make them pass without changing them, and without breaking any of the existing tests.

You are done when `go build ./...`, `go vet ./...` and `go test ./...` all succeed.
