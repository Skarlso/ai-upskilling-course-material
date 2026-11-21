# The toy repo

[Back to the course](../README.md)

*Generated*: This is a generated toy application that you can use for testing purposes.
It was a supervised generation process and contains a few things to look through.

Feel free to use it, or ignore it at your behest. It is there so group projects work on the SAME REPOSITORY! So you can compare results.

`pantry/` is a small Go program that keeps track of what is in a kitchen. This is for your agent to work on in [module 6](../modules/06-verification-part-1/README.md), [module 7](../modules/07-verification-part-2/README.md) and the [final project](../project/README.md).

## Structure

```
pantry/                 the Go program your agent works on
  cmd/pantry/           the command line
  internal/item/        what an item is, and how "500g flour" is parsed
  internal/store/       the items in memory and in a JSON file
  internal/report/      the table that `pantry list` prints
tasks/                  how a task looks, plus one example
```

## The program

```
$ pantry add 500g flour
Added. You now have 500g flour.
$ pantry add 3 eggs
Added. You now have 3 eggs.
$ pantry remove flour 200
Removed.
$ pantry list
ITEM   AMOUNT
eggs   3
flour  300g
2 items
```

Try it yourself: `cd pantry && go run ./cmd/pantry list`.

## On purpose things

- **Standard library only.** This is important for some of the modules so there is nothing to download.
- **The tests pass.** `go build ./... && go vet ./... && go test ./...` is green at the time of this writing. If something is red, please open an issue or check what you did.
- **It's small.** You can read it relatively quickly. Please do so before starting.
- **It has more than one package.** So the agent has to look in multiple places instead of just one.

## How to use it

Agents should work on their own copy:

1. Copy `pantry/` somewhere fresh.
2. Let the agent loose!
3. Check.
4. Delete.

A fresh copy every time is what makes multiple runs comparable. The things, and an examples you can run right now, are in [`tasks/`](tasks/README.md).

## Can I use my own repo instead?

Sure. Any small Go repo with tests will do. You just won't be able to compare your numbers with anybody else's.
