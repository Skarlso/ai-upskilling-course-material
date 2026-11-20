# Module 7: Verification, part 2

[Back to the course](../../README.md) | Previous: [Verification, part 1](../06-verification-part-1/README.md) | Next: [Security and sandboxing](../08-security-and-sandboxing/README.md)

In the last module we established a pass rate. In this module, we are going to be skeptical about it.

An agent that sees the tests is able to make the tests pass. Sometimes by actually fixing the thing that is broken, or... sometimes it will just delete the test. It's not being evil, it is just following orders which say "make the check pass". And it did that.

We can do a few things to make this work a bit more deterministic. There are checks that the agents just aren't aware of and then... there are your eyes!

## Read

The parts about scenarios and satisfaction in [The StrongDM Software Factory](https://factory.strongdm.ai/).

## Build

1. **Write three holdout scenarios.** These are checks that live outside of your agent's access. So it doesn't know about what it is checked against.
2. **Score against them.** After the agent finishes a task, run the scenarios on the result and record how many pass.
3. **Compare.** Put this score next to your pass rate from module 6. Mind the gap.
4. **Read the failures.** Open the transcript of every failed run and try grouping any failures that you find.

Step 4 IS important. Please do NOT skip it. It's important to understand problems and errors in order to understand the agent's thinking and working process, and to harden your harness.

## Done when

You know your agent's top two failure types and you understand the difference between what it can do and what it can't.

## Discuss

Where did the agent pass the check without doing the job?

## Extras

- **An LLM judge.** Label something that a command couldn't understand and then have another model judge the same failures and see if it agrees with you. If you don't label anything you aren't measuring.

## Optional reading

- [Evals for AI Engineers](https://www.oreilly.com/library/view/evals-for-ai/9798341660717/) by Hamel Husain and Shreya Shankar. This book is rather expensive, but it's worth it. It is quite dense.
