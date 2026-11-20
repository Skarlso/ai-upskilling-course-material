# Module 3: Context engineering

[Back to the course](../../README.md) | Previous: [The loop, part 2](../02-loop-part-2/README.md) | Next: [Harness engineering](../04-harness-engineering/README.md)

One important thing to understand is that a model will only ever know _what you send it_. Every time. Nothing else. Just what you send.

It doesn't know about your project, it doesn't know about conventions, or about what happened five minutes ago, unless... You guessed it. You send it with the text.
So the question "why did the agent do that?" is almost always answered by the same thing. Go and look at what it got!

That is _context engineering_. Deciding what goes in, what does not, and what gets thrown away. It's the only thing you can influence per call. So let's make sure you understand how it works.

## Read

Factor 3, "Own your context window", in [12-Factor Agents](https://github.com/humanlayer/12-factor-agents).

## Build

1. **Log every request, in full.** One file, one request per line, with everything: system prompt, messages, tool definitions and results! No summary! You need to be exact here.
2. **Count tokens per turn.** The API will tell you how many went in and out, and you need to keep a correct tally.
3. **Load an instructions file.** If there is an `AGENTS.md` in the working directory, put its contents into the system prompt. This is how every commercial coding agent learns about a project. And now you do too.
4. **Read a log end to end.** Really read it. This is the actual exercise.

## Done when

You can open the log and point to every piece of text the model saw on a given turn.

## Discuss

What surprised you when you read exactly what your agent sent to the model? Is it exciting, boring, surprising, horrifying, all of the above?

## Extras

- **Break it on purpose.** Make it read a huge file, or run a command with a wall of output, and watch the tokens burn. The answers will get worse and worse. Now try to fix it.
- **Compaction.** When the conversation gets close to the limit, replace the old turns with a summary. That's what compact does!
- **A subagent.** Start a second loop with a fresh, empty context, give it one job, and take back only its summary. Tada. Subagent.

## Optional reading

- "How coding agents work" in [Agentic Engineering Patterns](https://simonwillison.net/guides/agentic-engineering-patterns/) by Simon Willison
- Factor 9, on compacting errors, in 12-Factor Agents
