#!/bin/sh
# Checks the example-rename task.
#
#   check.sh <path-to-the-working-copy>
#
# Exits 0 when the task is done, non-zero otherwise.
set -eu

here=$(cd "$(dirname "$0")" && pwd)
work=${1:?usage: check.sh <path-to-the-working-copy>}

# Put the task's tests back first. If the agent edited or deleted them to get
# to green, that stops working right here.
cp -R "$here/setup/." "$work/"

cd "$work"
go build ./...
go vet ./...
go test ./...
