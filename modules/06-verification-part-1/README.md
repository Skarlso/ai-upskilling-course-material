# Module 6: Verification, part 1

[Back to the course](../../README.md) | Previous: [Tools and protocols](../05-tools-and-protocols/README.md) | Next: [Verification, part 2](../07-verification-part-2/README.md)

"It works on my agent."

We have all said it. With agents it means even less than usual, because the same agent on the same task can succeed on Monday and fail on Tuesday. Nothing changed. That is just what working with a model is like.

So we stop guessing and start _measuring_. This module and the next one are, in my opinion, the most important in the whole course. Writing code got cheap. Knowing if the code is right did not.

## Read

The red/green TDD chapter in the Testing and QA part of [Agentic Engineering Patterns](https://simonwillison.net/guides/agentic-engineering-patterns/) by Simon Willison.

## Build

You need a small Go repo with tests. Use the shared one in [`toy-repo/`](../../toy-repo/) so we can compare numbers.

1. **Define five tasks.** A task is two things: a prompt for the agent, and a check command that exits 0 when the job is done. `go build ./... && go test ./...` is a perfectly good check.
2. **Write a runner.** For each task, three times over: take a fresh copy of the repo, run your agent with the prompt, run the check, write down pass or fail.
3. **Print the pass rate.** Per task and overall.

What makes a good task? One where a command can tell you if it's done. Fix a failing test. Write a function that satisfies a test you provide. Add a flag to a CLI. Fix a bug described only in words.

## Done when

You can state a pass rate with the number of runs behind it. "11 out of 15", not "pretty good".

## Discuss

Which task was flaky? And was the task to blame, or the agent?

## A word on small numbers

Fifteen runs is not a lot. If one setup gets 11 and another gets 12, you have learned nothing yet. Big gaps mean something. Small ones need more runs.

## Extras

- Ten tasks, five runs each.
- Record tokens and cost per run. You will want them for the final project.
- Change _one_ line in your `AGENTS.md` and measure again. Congratulations, you can now tell if a change helped instead of hoping it did.

## Optional reading

- The free [evals email series](https://ai.hamel.dev/eval-course) by Hamel Husain and Shreya Shankar
