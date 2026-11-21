# Tasks

[Back to the toy repo](../README.md)

Agent tasks. Each with three things:

| | What it is |
| --- | --- |
| `prompt.md` | Prompt for the agent. Write it like an issue. |
| `setup/` | Copy of files in pantry. Optional. |
| `check.sh` | Takes the path to the working copy. Exits 0 when the job is done. |

Agent only ever sees its copy and never this folder!

## One run steps

1. Copy `pantry/` to a fresh, empty directory.
2. Copy the task's `setup/` over it, if it has one.
3. Start your agent in that directory and give it `prompt.md`.
4. When the agent says it's done, or hits its limits, run `check.sh` on the directory.
5. Exit code 0 means success. Anything else is a fail. Write that down.

By hand, from the root of the course repo, that looks like this:

```sh
work=$(mktemp -d)
cp -R toy-repo/pantry/. "$work/"
cp -R toy-repo/tasks/example-rename/setup/. "$work/"

# ... run your agent in $work with the contents of prompt.md ...

sh toy-repo/tasks/example-rename/check.sh "$work"
echo "exit code: $?"
```

Your runner in module 6 does exactly this, in a loop, and counts.

## The example

[`example-rename/`](example-rename/prompt.md) asks for a way to rename an item.

It has the red/green shape from the reading. `setup/` drops in tests that fail, because the feature doesn't exist yet. The agent's job is to make them pass.

Now, take a look at `check.sh`. Before it runs the tests, it copies the task's tests back in. Why does it do that? Because an agent that deletes a failing test has, technically, made the tests pass. And that's not gonna happen on our watch!

Try the check right now yourself and understand how it works. It should fail. That is your proof the task isn't already done.

## Writing your own

Write one yourself! That is half the lesson here. A good task has the following three things:

- **A command can tell if it's done.** No judgement calls yet, we'll build that in module 7.
- **It's small.** A few minutes of work for the agent, not an afternoon.
- **It fails before the agent starts.** Always run the check on a fresh copy first.

Some shapes that work well:

- Tests that fail, and a prompt that says "make them pass". Like the example.
- Features described in words, with the tests hidden in `check.sh` instead of `setup/`.
- Create a bug. `setup/` can overwrite a source file too, so you can plant one yourself.
- A refactor where the check is simply "everything still passes, and this function is gone".

And if you need ideas for `pantry`... sorting the list by amount, amounts like `1.5kg`, removing with a unit (`remove flour 200g`), a warning when something is running low, saving the file safely when the disk is full.

Important! Holdout scenarios from module 7 SHOULD NOT GO into the task's `setup/` because the whole point is that the agent never sees those! They are our independent judges.
