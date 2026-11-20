# Module 9: System design

[Back to the course](../../README.md) | Previous: [Security and sandboxing](../08-security-and-sandboxing/README.md) | Next: [Final project](../../project/README.md)

This is the last module! Welcome! You are doing fantastic so far. Now, we'll put everything to use.

An agent system is a distributed system with one very strange component in it. A part that is slow, costs money per call, and doesn't give the same answer when asked again. Everything you know about timeouts, retries and budgets applies. Even more than ever actually.

There is one question you have to be sure about when you do something. Does this thing _require_ an agent? A workflow is a lot cheaper and is deterministic, and in most cases, enough.

## Read

The workflow patterns section of [Building effective agents](https://www.anthropic.com/engineering/building-effective-agents) by Anthropic.

## Build

1. **Hard caps.** A maximum number of tokens, a maximum wall-clock time, and a maximum spend for one run. If any of them is reached, the agent needs to stop and tell you which one it hit.
2. **A human in the loop.** Any write outside the working directory blocks until user approval. Also, create a list of commands that need to be approved first before execution.

Small changes for sure, but they change how much you trust the thing.

## Done when

A run that hits its budget stops cleanly and tells you what it hit. A write outside the workspace is blocked for user action.

## Discuss

Do a design exercise on paper or on a whiteboard and talk about it. No code.

Design an agent that triages incoming issues for an open source project. It finds duplicate issues, or pull requests that might already address them. Workflow or agent? Where does the state live? What is it allowed to touch? What does one run cost? What happens when it's wrong? There are a lot of things to think about. How does it communicate its findings? Log, comment on the issue? With what account? Etc.

## Extras

- **Pause and resume.** Store each run as a log of events. Kill the process in the middle of a task. Start it again and have it continue from the log.
- **A second opinion.** A separate evaluator pass, with a fresh context, that reads the diff against the task.

## Optional reading

- Factors 5, 6, 10 and 12 in [12-Factor Agents](https://github.com/humanlayer/12-factor-agents)
