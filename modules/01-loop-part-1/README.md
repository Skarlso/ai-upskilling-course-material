# Module 1: Orientation and the loop, part 1

[Back to the course](../../README.md) | Next: [The loop, part 2](../02-loop-part-2/README.md)

Let's start by taking the magic away.

An agent is not a mysterious thing. Simon Willison's definition reads as follows: "Agents run tools in a loop to achieve a goal." That's it. A model on a loop with a bunch of tools at its disposal. In this module you write the first half of that with your bare hands, and after that, coding agents will make sense to you.

And yes, _by hand_. No agents for this one. I know. But this is the one place in the course where typing it out yourself is the whole point!

## Read

[The Five Levels](https://simonwillison.net/2026/Jan/28/the-five-levels/). A short summary of Dan Shapiro's ladder, from autocomplete all the way up to a factory. Where are _you_ on this list?

## Build

Follow the first half of [How to Build an Agent](https://ampcode.com/how-to-build-an-agent) by Thorsten Ball. It's in Go, and it's an absolute joy to read.

1. **The chat loop.** Read a line from the terminal, send the whole conversation to the model, print the reply, repeat.
2. **A tool to read a file.** Provide a name, description and schema. Then, when the model asks, run it and send back the result.
3. **A tool to list files.** Same thing again. Now the model can look around.

Two things to notice while you do this:

- The model remembers _nothing_. You send the entire conversation _always_. The "memory" is nothing but a slice in your program.
- The model never runs anything. It only _asks_. After that, it's your code's job to decide what happens next.

The post uses a provider API but any model with tool calling will do, even a local one.

## Done when

You ask your agent a question about a file in a repo, and it goes and reads the file to answer you. That's it.

## Discuss

Which of the five levels do you work at today? What stops you from moving up one? And do you even want to?

## Extras

- Print every tool call as it happens, with arguments. You will want this for the rest of the course.
- Ask it something that will need two files and see how it will chain calls on its own.

## Optional reading

- [What is agentic engineering?](https://simonwillison.net/guides/agentic-engineering-patterns/what-is-agentic-engineering/) by Simon Willison
- [Building effective agents](https://www.anthropic.com/engineering/building-effective-agents) by Anthropic
