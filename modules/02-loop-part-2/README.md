# Module 2: The loop, part 2

[Back to the course](../../README.md) | Previous: [The loop, part 1](../01-loop-part-1/README.md) | Next: [Context engineering](../03-context-engineering/README.md)

So far we built eyes for the agent. Now let's add limbs.

This is the module where it stops being a chatbot and becomes a _coding agent_. It will edit files. It will run commands. And the first time it reads a test failure, fixes the code and runs the test again without you saying a word... that's when you WILL feel the power of the sun in the palm of your hands.

It's also the moment it can do damage! Like Otto Octavius, you'd better have a backup plan, and not just give it access to your entire home folder.

## Read

The second half of [How to Build an Agent](https://ampcode.com/how-to-build-an-agent).

## Build

1. **A tool to edit files.** The post does it with a simple string replace. That is enough. Resist the urge to do more right now, otherwise, you'll get overwhelmed quickly.
2. **A tool to run shell commands.** Return stdout, stderr and the exit code. Give it a timeout because commands can hang indefinitely.
3. **Put the model behind a Go interface.** One small interface for "send these messages and tools, get a reply". Make sure your loop doesn't depend on any provider.
4. **Give it a job.** Make a tiny Go repo with one failing test and then sic your agent on it and tell it to make it work. (Hopefully it won't just delete the test...)

## Done when

Your agent fixes a failing test in a small Go repo without your help.

## Discuss

What did the model do that you did not expect, once it could run commands? This one can result in some very interesting conversations.

## Extras

- Use the interface you just wrote: run the same agent against a second provider, or a local model, and watch for differences.
- Cut off very long tool output before you send it back. You will see why in the next module.
- [how-to-build-a-coding-agent](https://github.com/ghuntley/how-to-build-a-coding-agent) by Geoffrey Huntley is a second Go walkthrough of the same idea, in six steps. Good if you want another angle.

## A note before you move on

Keep this agent. You will take it through every module after this one. It doesn't need to be pretty! But it needs to be _yours_.
